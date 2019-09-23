package sonar

import (
	"context"
	"fmt"
	"github.com/dchest/uniuri"
	jenkinsHelper "github.com/epmd-edp/jenkins-operator/v2/pkg/controller/jenkinsscript/helper"
	keycloakApi "github.com/epmd-edp/keycloak-operator/pkg/apis/v1/v1alpha1"
	"github.com/epmd-edp/sonar-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/sonar-operator/v2/pkg/client/sonar"
	pkgHelper "github.com/epmd-edp/sonar-operator/v2/pkg/helper"
	"github.com/epmd-edp/sonar-operator/v2/pkg/service/platform"
	sonarHelper "github.com/epmd-edp/sonar-operator/v2/pkg/service/sonar/helper"
	sonarSpec "github.com/epmd-edp/sonar-operator/v2/pkg/service/sonar/spec"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
	k8sErr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	JenkinsLogin            = "jenkins"
	JenkinsUsername         = "Jenkins"
	ReaduserLogin           = "read"
	ReaduserUsername        = "Read-only user"
	NonInteractiveGroupName = "non-interactive-users"
	WebhookUrl              = "http://jenkins:8080/sonarqube-webhook/"
	DefaultPassword         = "admin"
	ClaimName               = "roles"
	SonarPort               = 9000

	defaultConfigFilesAbsolutePath = "/usr/local/"
	localConfigsRelativePath       = "configs"

	defaultProfileAbsolutePath = defaultConfigFilesAbsolutePath + localConfigsRelativePath + "/" + defaultQualityProfilesFileName

	defaultQualityProfilesFileName = "quality-profile.xml"
)

type Client struct {
	client resty.Client
}

type SonarService interface {
	// This is an entry point for service package. Invoked in err = r.service.Install(*instance) sonar_controller.go, Reconcile method.
	Install(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error)
	Configure(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error, bool)
	ExposeConfiguration(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error)
	Integration(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error)
	IsDeploymentConfigReady(instance v1alpha1.Sonar) (bool, error)
}

func NewSonarService(platformService platform.PlatformService, k8sClient client.Client, k8sScheme *runtime.Scheme) SonarService {
	return SonarServiceImpl{platformService: platformService, k8sClient: k8sClient, k8sScheme: k8sScheme}
}

type SonarServiceImpl struct {
	// Providing sonar service implementation through the interface (platform abstract)
	platformService platform.PlatformService
	k8sClient       client.Client
	k8sScheme       *runtime.Scheme
}

func (s SonarServiceImpl) initSonarClient(instance *v1alpha1.Sonar, defaultPassword bool) (*sonar.SonarClient, error) {
	sonarRoute, scheme, err := s.platformService.GetRoute(instance.Namespace, instance.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get route for %v", instance.Name)
	}
	sonarApiUrl := fmt.Sprintf("%v://%v/api", scheme, sonarRoute.Spec.Host)
	sc := &sonar.SonarClient{}

	password := DefaultPassword
	if !defaultPassword {
		adminSecretName := fmt.Sprintf("%v-admin-password", instance.Name)
		credentials, err := s.platformService.GetSecretData(instance.Namespace, adminSecretName)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to get %v secret for Sonar client!", adminSecretName)
		}
		password = string(credentials["password"])
	}

	err = sc.InitNewRestClient(sonarApiUrl, "admin", password)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to init new Sonar client!")
	}
	return sc, nil
}

func (s SonarServiceImpl) Integration(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error) {
	sc, err := s.initSonarClient(&instance, false)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to initialize Sonar Client!")
	}
	realm, err := s.getKeycloakRealm(instance)
	if err != nil {
		return &instance, err
	}
	if realm != nil {
		if realm.Annotations == nil {
			return &instance, errors.New("realm main does not have required annotations")
		}
		openIdConfiguration := realm.Annotations["openid-configuration"]
		err = sc.ConfigureGeneralSettings("value", "sonar.auth.oidc.providerConfiguration", openIdConfiguration)
		if err != nil {
			return &instance, errors.Wrap(err, "Failed to to configure sonar.auth.oidc.providerConfiguration!")
		}
	}
	sonarRoute, scheme, err := s.platformService.GetRoute(instance.Namespace, instance.Name)
	var baseUrl string
	if sonarRoute != nil {
		baseUrl = fmt.Sprintf("%v://%v", scheme, sonarRoute.Spec.Host)
		err = sc.ConfigureGeneralSettings("value", "sonar.core.serverBaseURL", baseUrl)
		if err != nil {
			return &instance, errors.Wrap(err, "Failed to configure sonar.core.serverBaseURL!")
		}
	}
	cl, err := s.getKeycloakClient(instance)
	if err != nil {
		return &instance, err
	}

	if cl == nil {
		err = s.createKeycloakClient(instance, baseUrl)
	}

	if err != nil {
		return &instance, err
	}

	err = sc.ConfigureGeneralSettings("value", "sonar.auth.oidc.clientId.secured", instance.Name)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to configure sonar.auth.oidc.clientId.secured!")
	}

	err = sc.ConfigureGeneralSettings("value", "sonar.auth.oidc.groupsSync.claimName", ClaimName)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to configure sonar.auth.oidc.groupsSync.claimName!")
	}

	err = sc.ConfigureGeneralSettings("value", "sonar.auth.oidc.groupsSync", "true")
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to configure sonar.auth.oidc.groupsSync!")
	}

	err = sc.ConfigureGeneralSettings("value", "sonar.auth.oidc.enabled", "true")
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to configure sonar.auth.oidc.enabled!")
	}
	return &instance, nil
}

func (s SonarServiceImpl) getKeycloakRealm(instance v1alpha1.Sonar) (*keycloakApi.KeycloakRealm, error) {
	realm := &keycloakApi.KeycloakRealm{}
	err := s.k8sClient.Get(context.TODO(), types.NamespacedName{
		Name:      "main",
		Namespace: instance.Namespace,
	}, realm)
	if err != nil {
		if k8sErr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return realm, nil
}

func (s SonarServiceImpl) getKeycloakClient(instance v1alpha1.Sonar) (*keycloakApi.KeycloakClient, error) {
	cl := &keycloakApi.KeycloakClient{}
	err := s.k8sClient.Get(context.TODO(), types.NamespacedName{
		Name:      instance.Name,
		Namespace: instance.Namespace,
	}, cl)
	if err != nil {
		if k8sErr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return cl, nil
}

func (s SonarServiceImpl) createKeycloakClient(instance v1alpha1.Sonar, baseUrl string) error {
	cl := &keycloakApi.KeycloakClient{
		ObjectMeta: v1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: keycloakApi.KeycloakClientSpec{
			ClientId:                instance.Name,
			Public:                  true,
			WebUrl:                  baseUrl,
			AdvancedProtocolMappers: true,
			RealmRoles: &[]keycloakApi.RealmRole{
				{
					Name:      "sonar-administrators",
					Composite: "administrator",
				},
				{
					Name:      "sonar-users",
					Composite: "developer",
				},
			},
		},
	}
	return s.k8sClient.Create(context.TODO(), cl)
}

func (s SonarServiceImpl) ExposeConfiguration(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error) {

	externalConfig := v1alpha1.SonarExternalConfiguration{nil, nil, nil}

	sc, err := s.initSonarClient(&instance, false)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to initialize Sonar Client!")
	}

	jenkinsPassword := uniuri.New()
	err = sc.CreateUser(JenkinsLogin, JenkinsUsername, jenkinsPassword)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create user %v in Sonar!", JenkinsUsername)
	}

	err = sc.AddUserToGroup(NonInteractiveGroupName, JenkinsLogin)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to add %v user in %v group!", JenkinsLogin, NonInteractiveGroupName)
	}

	err = sc.AddPermissionsToUser(JenkinsLogin, "admin")
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to add admin persmissions to  %v user", JenkinsLogin)
	}

	ciToken, err := sc.GenerateUserToken(JenkinsLogin)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to generate token for %v user", JenkinsLogin)
	}

	ciUserName := fmt.Sprintf("%v-ciuser-token", instance.Name)
	if ciToken != nil {
		ciSecret := map[string][]byte{
			"username": []byte(JenkinsLogin),
			"secret":   []byte(*ciToken),
		}

		err = s.platformService.CreateSecret(instance, ciUserName, ciSecret)
		if err != nil {
			return &instance, errors.Wrapf(err, "Failed to create secret for  %v user", ciUserName)
		}
	}

	err = s.platformService.CreateJenkinsServiceAccount(instance.Namespace, ciUserName, "token")
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Jenkins Service Account for %v", ciUserName)
	}

	data := sonarHelper.InitNewJenkinsPluginInfo(true)
	data.ServerName = instance.Name
	data.SecretName = JenkinsLogin

	jenkinsScriptContext, err := sonarHelper.ParseDefaultTemplate(data)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to parse default Jenkins plugin template!")
	}

	configMapName := fmt.Sprintf("%s-%s", instance.Name, sonarSpec.JenkinsPluginConfigPostfix)
	configMapData := map[string]string{
		jenkinsHelper.JenkinsDefaultScriptConfigMapKey: jenkinsScriptContext.String(),
	}

	err = s.platformService.CreateConfigMap(instance, configMapName, configMapData)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Config Map %v", configMapName)
	}

	err = s.platformService.CreateJenkinsScript(instance.Namespace, configMapName)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Jenkins Script for %v", ciUserName)
	}

	readPassword := uniuri.New()
	err = sc.CreateUser(ReaduserLogin, ReaduserUsername, readPassword)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create %v user in Sonar!", ReaduserUsername)
	}

	readToken, err := sc.GenerateUserToken(ReaduserLogin)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to generate token for %v user", ReaduserLogin)
	}

	readUserSecretName := fmt.Sprintf("%v-readuser-token", instance.Name)
	if readToken != nil {
		readSecret := map[string][]byte{
			"username": []byte(ReaduserLogin),
			"token":    []byte(*readToken),
		}

		err = s.platformService.CreateSecret(instance, readUserSecretName, readSecret)
		if err != nil {
			return &instance, errors.Wrapf(err, "Failed to create secret for  %v user", readUserSecretName)
		}
	}

	err = sc.AddUserToGroup(NonInteractiveGroupName, ReaduserLogin)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to add %v user in %v group!", ReaduserLogin, NonInteractiveGroupName)
	}

	identityServerSecretName := fmt.Sprintf("%v-is-credentials", instance.Name)
	identityServiceClientSecret := uniuri.New()
	identityServiceClientCredenrials := map[string][]byte{
		"client_id":     []byte(instance.Name),
		"client_secret": []byte(identityServiceClientSecret),
	}

	err = s.platformService.CreateSecret(instance, identityServerSecretName, identityServiceClientCredenrials)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create secret for  %v Keycloak client!", readUserSecretName)
	}

	externalConfig.AdminUser = &v1alpha1.SonarExternalConfigurationItem{instance.Name + "-admin-password", "Secret", "Password for Sonar admin user"}
	externalConfig.ReadUser = &v1alpha1.SonarExternalConfigurationItem{instance.Name + "-readuser-token", "Secret", "Token for read-only user"}
	externalConfig.IsCredentials = &v1alpha1.SonarExternalConfigurationItem{instance.Name + "-is-credentials", "Secret", "Credentials for Identity Server integration"}

	err = s.updateExternalConfig(&instance, externalConfig)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to update ExternalConfig field in Sonar spec!")
	}

	return &instance, nil
}

func (s SonarServiceImpl) Configure(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error, bool) {
	sc, err := s.initSonarClient(&instance, true)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to initialize Sonar Client!"), false
	}

	// TODO(Serhii Shydlovskyi): Error handling here ?
	sc.WaitForStatusIsUp(60, 10)

	adminSecretName := fmt.Sprintf("%v-admin-password", instance.Name)
	credentials, err := s.platformService.GetSecretData(instance.Namespace, adminSecretName)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to get secret data from %v!", adminSecretName), false
	}
	password := string(credentials["password"])
	// TODO(Serhii Shydlovskyi): Add check for password presence. Breaks status update.
	sc.ChangePassword("admin", DefaultPassword, password)

	sc, err = s.initSonarClient(&instance, false)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to initialize Sonar Client!"), false
	}

	plugins := []string{"authoidc", "checkstyle", "findbugs", "pmd", "jacoco", "xml", "javascript"}
	err = sc.InstallPlugins(plugins)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to install plugins for Sonar!"), false
	}

	executableFilePath := pkgHelper.GetExecutableFilePath()
	profilePath := defaultProfileAbsolutePath

	if _, err := k8sutil.GetOperatorNamespace(); err != nil && err == k8sutil.ErrNoNamespace {
		profilePath = fmt.Sprintf("%v\\..\\%v\\%v", executableFilePath, localConfigsRelativePath, defaultQualityProfilesFileName)
	}
	_, err = sc.UploadProfile("EDP way", profilePath)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to upload EDP way profile!"), false
	}

	qgContidions := []map[string]string{
		{"error": "80", "metric": "new_coverage", "op": "LT", "period": "1"},
		{"error": "0", "metric": "test_errors", "op": "GT"},
		{"error": "3", "metric": "new_duplicated_lines_density", "op": "GT", "period": "1"},
		{"error": "0", "metric": "test_failures", "op": "GT"},
		{"error": "0", "metric": "blocker_violations", "op": "GT"},
		{"error": "0", "metric": "critical_violations", "op": "GT"},
	}
	_, err = sc.CreateQualityGate("EDP way", qgContidions)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to configure EDP way quality gate!"), false
	}

	err = sc.CreateGroup(NonInteractiveGroupName)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create %v group!", NonInteractiveGroupName), false
	}

	err = sc.AddPermissionsToGroup(NonInteractiveGroupName, "scan")
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to add scan permission for %v group!", NonInteractiveGroupName), false
	}

	err = sc.AddWebhook(JenkinsLogin, WebhookUrl)
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to add Jenkins webhook!"), false
	}

	err = sc.ConfigureGeneralSettings("values", "sonar.typescript.lcov.reportPaths", "coverage/lcov.info")
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to configure sonar.typescript.lcov.reportPaths!"), false
	}

	err = sc.ConfigureGeneralSettings("values", "sonar.coverage.jacoco.xmlReportPaths", "target/site/jacoco/jacoco.xml")
	if err != nil {
		return &instance, errors.Wrap(err, "Failed to configure sonar.coverage.jacoco.xmlReportPaths!"), false
	}

	return &instance, nil, true
}

// Invoking install method against SonarServiceImpl object should trigger list of methods, stored in client edp.PlatformService
func (s SonarServiceImpl) Install(instance v1alpha1.Sonar) (*v1alpha1.Sonar, error) {

	dbSecret := map[string][]byte{
		"database-user":     []byte("admin"),
		"database-password": []byte(uniuri.New()),
	}

	sonarDbName := fmt.Sprintf("%v-db", instance.Name)
	err := s.platformService.CreateSecret(instance, sonarDbName, dbSecret)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create secret for %s", sonarDbName)
	}

	adminSecret := map[string][]byte{
		"user":     []byte("admin"),
		"password": []byte(uniuri.New()),
	}

	adminSecretName := fmt.Sprintf("%v-admin-password", instance.Name)
	err = s.platformService.CreateSecret(instance, adminSecretName, adminSecret)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create password for Admin in %s Sonar!", instance.Name)
	}

	sa, err := s.platformService.CreateServiceAccount(instance)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Service Account for %v Sonar!", instance.Name)
	}

	err = s.platformService.CreateSecurityContext(instance, sa)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Security Context for %v Sonar!", instance.Name)
	}

	err = s.platformService.CreateDeployConf(instance)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Deployment Config for %v Sonar!", instance.Name)
	}

	err = s.platformService.CreateService(instance)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Service for %v Sonar!", instance.Name)
	}

	err = s.platformService.CreateExternalEndpoint(instance)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Route for %v Sonar!", instance.Name)
	}

	err = s.platformService.CreateVolume(instance)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create Volume for %v Sonar!", instance.Name)
	}

	err = s.platformService.CreateDbDeployConf(instance)
	if err != nil {
		return &instance, errors.Wrapf(err, "Failed to create database Deploymetn Config for %v Sonar!", instance.Name)
	}

	return &instance, nil
}

func (s SonarServiceImpl) updateExternalConfig(instance *v1alpha1.Sonar, config v1alpha1.SonarExternalConfiguration) error {
	instance.Spec.SonarExternalConfiguration = config

	err := s.k8sClient.Status().Update(context.TODO(), instance)
	if err != nil {
		err = s.k8sClient.Update(context.TODO(), instance)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s SonarServiceImpl) IsDeploymentConfigReady(instance v1alpha1.Sonar) (bool, error) {
	sonarIsReady := false

	sonarDc, err := s.platformService.GetDeploymentConfig(instance)
	if err != nil {
		return sonarIsReady, err
	}

	if sonarDc.Status.AvailableReplicas == 1 {
		sonarIsReady = true
	}

	return sonarIsReady, nil
}
