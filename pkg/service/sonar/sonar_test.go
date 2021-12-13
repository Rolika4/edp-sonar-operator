package sonar

import (
	"context"
	"fmt"
	"testing"
	"time"

	jenkinsV1Api "github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1alpha1"

	sonarClient "github.com/epam/edp-sonar-operator/v2/pkg/client/sonar"

	"github.com/epam/edp-sonar-operator/v2/pkg/service/platform"

	"github.com/epam/edp-sonar-operator/v2/pkg/apis/edp/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	coreV1Api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSonarServiceImpl_DeleteResource(t *testing.T) {
	secret := coreV1Api.Secret{ObjectMeta: metav1.ObjectMeta{Name: "name", Namespace: "ns"}}
	s := Service{
		k8sClient: fake.NewClientBuilder().WithRuntimeObjects(&secret).Build(),
	}

	if _, err := s.DeleteResource(context.Background(), &secret, "fin", func() error {
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	secret.DeletionTimestamp = &metav1.Time{Time: time.Now()}
	secret.Finalizers = []string{"fin"}
	s.k8sClient = fake.NewClientBuilder().WithRuntimeObjects(&secret).Build()

	if _, err := s.DeleteResource(context.Background(), &secret, "fin", func() error {
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestServiceMock_Configure(t *testing.T) {
	sch := runtime.NewScheme()
	if err := v1alpha1.AddToScheme(sch); err != nil {
		t.Fatal(err)
	}
	if err := coreV1Api.AddToScheme(sch); err != nil {
		t.Fatal(err)
	}
	if err := jenkinsV1Api.AddToScheme(sch); err != nil {
		t.Fatal(err)
	}

	snr := v1alpha1.Sonar{ObjectMeta: metav1.ObjectMeta{
		Namespace: "ns", Name: "snr1",
	}, Spec: v1alpha1.SonarSpec{DefaultPermissionTemplate: "tpl123"}}

	jns := jenkinsV1Api.Jenkins{Spec: jenkinsV1Api.JenkinsSpec{BasePath: "zabagdo"}, ObjectMeta: metav1.ObjectMeta{
		Name: "js1", Namespace: snr.Namespace,
	}}
	plMock := platform.Mock{}
	clMock := ClientMock{}

	s := Service{
		k8sClient:       fake.NewClientBuilder().WithScheme(sch).WithRuntimeObjects(&jns).Build(),
		platformService: &plMock,
		sonarClientBuilder: func(ctx context.Context, instance *v1alpha1.Sonar,
			useDefaultPassword bool) (ClientInterface, error) {
			return &clMock, nil
		},
	}

	adminSecret := coreV1Api.Secret{Data: map[string][]byte{
		"password": []byte("pwd123"),
	}}

	plMock.On("CreateSecret", snr.Name, snr.Namespace,
		fmt.Sprintf("%s-admin-password", snr.Name)).Return(&adminSecret, nil)
	plMock.On("SetOwnerReference", &snr, &adminSecret).Return(nil)
	plMock.On("GetExternalEndpoint", snr.Namespace, snr.Name).Return("url", nil)
	clMock.On("ChangePassword", "admin", "admin", "pwd123").Return(nil)
	clMock.On("InstallPlugins",
		[]string{"authoidc", "checkstyle", "findbugs", "pmd", "jacoco", "xml", "javascript", "go", "ansible", "yaml",
			"python", "csharp", "groovy"}).Return(nil)
	clMock.On("UploadProfile", "EDP way").
		Return("profile123", nil)
	clMock.On("CreateQualityGate", "EDP way").Return("qg1", nil)
	clMock.On("GetGroup", nonInteractiveGroupName).Return(nil, sonarClient.ErrNotFound("not found"))
	clMock.On("CreateGroup", &sonarClient.Group{Name: nonInteractiveGroupName}).Return(nil)
	clMock.On("AddPermissionsToGroup", nonInteractiveGroupName, "scan").Return(nil)
	clMock.On("AddWebhook", "jenkins",
		"http://jenkins.ns:8080/zabagdo/sonarqube-webhook/").Return(nil)
	clMock.On("ConfigureGeneralSettings", "values", "sonar.typescript.lcov.reportPaths",
		"coverage/lcov.info").Return(nil)
	clMock.On("ConfigureGeneralSettings", "values", "sonar.coverage.jacoco.xmlReportPaths",
		"target/site/jacoco/jacoco.xml").Return(nil)
	clMock.On("SetDefaultPermissionTemplate", snr.Spec.DefaultPermissionTemplate).Return(nil)

	if err := s.Configure(context.Background(), &snr); err != nil {
		t.Fatalf("%+v", err)
	}
}