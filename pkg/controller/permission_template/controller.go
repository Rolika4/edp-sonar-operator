package permission_template

import (
	"context"
	"reflect"

	sonarApi "github.com/epam/edp-sonar-operator/v2/pkg/apis/edp/v1alpha1"
	sonarClient "github.com/epam/edp-sonar-operator/v2/pkg/client/sonar"
	"github.com/epam/edp-sonar-operator/v2/pkg/helper"
	"github.com/epam/edp-sonar-operator/v2/pkg/service/platform"
	"github.com/epam/edp-sonar-operator/v2/pkg/service/sonar"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const finalizer = "sonar.permission_template.operator"

type Reconcile struct {
	service sonar.SonarService
	client  client.Client
	log     logr.Logger
}

func NewReconcile(client client.Client, scheme *runtime.Scheme, log logr.Logger, platformType string) (*Reconcile, error) {
	ps, err := platform.NewPlatformService(platformType, scheme, client)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create platform service")
	}

	return &Reconcile{
		service: sonar.NewSonarService(ps, client, scheme),
		client:  client,
		log:     log.WithName("permission-template"),
	}, nil
}

func (r *Reconcile) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sonarApi.SonarPermissionTemplate{}, builder.WithPredicates(predicate.Funcs{
			UpdateFunc: isSpecUpdated,
		})).
		Complete(r)
}

func isSpecUpdated(e event.UpdateEvent) bool {
	oo := e.ObjectOld.(*sonarApi.SonarPermissionTemplate)
	no := e.ObjectNew.(*sonarApi.SonarPermissionTemplate)

	return !reflect.DeepEqual(oo.Spec, no.Spec) ||
		(oo.GetDeletionTimestamp().IsZero() && !no.GetDeletionTimestamp().IsZero())
}

func (r *Reconcile) Reconcile(ctx context.Context, request reconcile.Request) (result reconcile.Result, resultErr error) {
	log := r.log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	log.Info("Reconciling SonarPermissionTemplate")

	var instance sonarApi.SonarPermissionTemplate
	if err := r.client.Get(ctx, request.NamespacedName, &instance); err != nil {
		if k8sErrors.IsNotFound(err) {
			log.Info("instance not found")
			return
		}

		resultErr = errors.Wrap(err, "unable to get sonar permission tpl from k8s")
		return
	}

	if err := r.tryReconcile(ctx, &instance); err != nil {
		instance.Status.Value = err.Error()
		result.RequeueAfter = helper.SetFailureCount(&instance)
		log.Error(err, "an error has occurred while handling keycloak realm idp", "name",
			request.Name)
	} else {
		helper.SetSuccessStatus(&instance)
	}

	if err := r.client.Status().Update(ctx, &instance); err != nil {
		resultErr = errors.Wrap(err, "unable to update status")
	}

	log.Info("Reconciling done")
	return
}

func (r *Reconcile) tryReconcile(ctx context.Context, instance *sonarApi.SonarPermissionTemplate) error {
	sClient, err := r.service.ClientForChild(ctx, instance)
	if err != nil {
		return errors.Wrap(err, "unable to init sonar rest client")
	}

	_, err = sClient.GetPermissionTemplate(ctx, instance.Spec.Name)
	if sonarClient.IsErrNotFound(err) {
		sonarPermTpl := specToClientTemplate(&instance.Spec)

		if err := sClient.CreatePermissionTemplate(ctx, sonarPermTpl); err != nil {
			return errors.Wrap(err, "unable to create sonar permission template")
		}
		instance.Status.ID = sonarPermTpl.ID
	} else if err != nil {
		return errors.Wrap(err, "unexpected error during get sonar permission template")
	} else {
		if instance.Status.ID == "" {
			return errors.New("permission template already exists in sonar")
		}

		tpl := specToClientTemplate(&instance.Spec)
		tpl.ID = instance.Status.ID

		if err := sClient.UpdatePermissionTemplate(ctx, tpl); err != nil {
			return errors.Wrap(err, "unable to update group")
		}
	}

	if err := syncPermissionTemplateGroups(ctx, instance, sClient); err != nil {
		return errors.Wrap(err, "unable to sync permission template groups")
	}

	if _, err := r.service.DeleteResource(ctx, instance, finalizer, func() error {
		if err := sClient.DeletePermissionTemplate(ctx, instance.Status.ID); err != nil {
			return errors.Wrap(err, "unable to delete permission template")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "unable to delete resource")
	}

	return nil
}

func specToClientTemplate(spec *sonarApi.SonarPermissionTemplateSpec) *sonarClient.PermissionTemplate {
	return &sonarClient.PermissionTemplate{Name: spec.Name, Description: spec.Description,
		ProjectKeyPattern: spec.ProjectKeyPattern}
}

func syncPermissionTemplateGroups(ctx context.Context, instance *sonarApi.SonarPermissionTemplate,
	sClient sonar.ClientInterface) error {
	groups, err := sClient.GetPermissionTemplateGroups(ctx, instance.Status.ID)
	if err != nil {
		return errors.Wrap(err, "unable to get permission template groups")
	}

	for _, g := range groups {
		if err := sClient.RemoveGroupFromPermissionTemplate(ctx, &g); err != nil {
			return errors.Wrap(err, "unable to remote group from permission template")
		}
	}

	for _, g := range instance.Spec.GroupPermissions {
		if err := sClient.AddGroupToPermissionTemplate(ctx, &sonarClient.PermissionTemplateGroup{
			GroupName:   g.GroupName,
			TemplateID:  instance.Status.ID,
			Permissions: g.Permissions,
		}); err != nil {
			return errors.Wrap(err, "unable to add group to permission template")
		}
	}

	return nil
}
