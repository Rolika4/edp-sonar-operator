// Code generated by mockery v2.9.4. DO NOT EDIT.

package mock

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	mock "github.com/stretchr/testify/mock"

	sonar "github.com/epam/edp-sonar-operator/v2/pkg/service/sonar"

	v1alpha1 "github.com/epam/edp-sonar-operator/v2/pkg/apis/edp/v1alpha1"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// ClientForChild provides a mock function with given fields: ctx, instance
func (_m *ServiceInterface) ClientForChild(ctx context.Context, instance sonar.ChildInstance) (sonar.ClientInterface, error) {
	ret := _m.Called(ctx, instance)

	var r0 sonar.ClientInterface
	if rf, ok := ret.Get(0).(func(context.Context, sonar.ChildInstance) sonar.ClientInterface); ok {
		r0 = rf(ctx, instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sonar.ClientInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, sonar.ChildInstance) error); ok {
		r1 = rf(ctx, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Configure provides a mock function with given fields: ctx, instance
func (_m *ServiceInterface) Configure(ctx context.Context, instance *v1alpha1.Sonar) error {
	ret := _m.Called(ctx, instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Sonar) error); ok {
		r0 = rf(ctx, instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteResource provides a mock function with given fields: ctx, instance, finalizer, deleteFunc
func (_m *ServiceInterface) DeleteResource(ctx context.Context, instance sonar.Deletable, finalizer string, deleteFunc func() error) (bool, error) {
	ret := _m.Called(ctx, instance, finalizer, deleteFunc)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, sonar.Deletable, string, func() error) bool); ok {
		r0 = rf(ctx, instance, finalizer, deleteFunc)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, sonar.Deletable, string, func() error) error); ok {
		r1 = rf(ctx, instance, finalizer, deleteFunc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExposeConfiguration provides a mock function with given fields: ctx, instance
func (_m *ServiceInterface) ExposeConfiguration(ctx context.Context, instance *v1alpha1.Sonar) error {
	ret := _m.Called(ctx, instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Sonar) error); ok {
		r0 = rf(ctx, instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Integration provides a mock function with given fields: ctx, instance
func (_m *ServiceInterface) Integration(ctx context.Context, instance *v1alpha1.Sonar) (*v1alpha1.Sonar, error) {
	ret := _m.Called(ctx, instance)

	var r0 *v1alpha1.Sonar
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Sonar) *v1alpha1.Sonar); ok {
		r0 = rf(ctx, instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Sonar)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.Sonar) error); ok {
		r1 = rf(ctx, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsDeploymentReady provides a mock function with given fields: instance
func (_m *ServiceInterface) IsDeploymentReady(instance *v1alpha1.Sonar) (bool, error) {
	ret := _m.Called(instance)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*v1alpha1.Sonar) bool); ok {
		r0 = rf(instance)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.Sonar) error); ok {
		r1 = rf(instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// K8sClient provides a mock function with given fields:
func (_m *ServiceInterface) K8sClient() client.Client {
	ret := _m.Called()

	var r0 client.Client
	if rf, ok := ret.Get(0).(func() client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	return r0
}
