// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import (
	appsv1 "k8s.io/api/apps/v1"
	apiautoscalingv1 "k8s.io/api/autoscaling/v1"

	autoscalingv1 "k8s.io/client-go/applyconfigurations/autoscaling/v1"

	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/apps/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// DeploymentInterface is an autogenerated mock type for the DeploymentInterface type
type DeploymentInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, deployment, opts
func (_m *DeploymentInterface) Apply(ctx context.Context, deployment *v1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.Deployment, error) {
	ret := _m.Called(ctx, deployment, opts)

	var r0 *appsv1.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *v1.DeploymentApplyConfiguration, metav1.ApplyOptions) *appsv1.Deployment); ok {
		r0 = rf(ctx, deployment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.DeploymentApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, deployment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyScale provides a mock function with given fields: ctx, deploymentName, scale, opts
func (_m *DeploymentInterface) ApplyScale(ctx context.Context, deploymentName string, scale *autoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (*apiautoscalingv1.Scale, error) {
	ret := _m.Called(ctx, deploymentName, scale, opts)

	var r0 *apiautoscalingv1.Scale
	if rf, ok := ret.Get(0).(func(context.Context, string, *autoscalingv1.ScaleApplyConfiguration, metav1.ApplyOptions) *apiautoscalingv1.Scale); ok {
		r0 = rf(ctx, deploymentName, scale, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiautoscalingv1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *autoscalingv1.ScaleApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, deploymentName, scale, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyStatus provides a mock function with given fields: ctx, deployment, opts
func (_m *DeploymentInterface) ApplyStatus(ctx context.Context, deployment *v1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.Deployment, error) {
	ret := _m.Called(ctx, deployment, opts)

	var r0 *appsv1.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *v1.DeploymentApplyConfiguration, metav1.ApplyOptions) *appsv1.Deployment); ok {
		r0 = rf(ctx, deployment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.DeploymentApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, deployment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, deployment, opts
func (_m *DeploymentInterface) Create(ctx context.Context, deployment *appsv1.Deployment, opts metav1.CreateOptions) (*appsv1.Deployment, error) {
	ret := _m.Called(ctx, deployment, opts)

	var r0 *appsv1.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.Deployment, metav1.CreateOptions) *appsv1.Deployment); ok {
		r0 = rf(ctx, deployment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.Deployment, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, deployment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *DeploymentInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *DeploymentInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *DeploymentInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*appsv1.Deployment, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *appsv1.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *appsv1.Deployment); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScale provides a mock function with given fields: ctx, deploymentName, options
func (_m *DeploymentInterface) GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (*apiautoscalingv1.Scale, error) {
	ret := _m.Called(ctx, deploymentName, options)

	var r0 *apiautoscalingv1.Scale
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *apiautoscalingv1.Scale); ok {
		r0 = rf(ctx, deploymentName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiautoscalingv1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, deploymentName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *DeploymentInterface) List(ctx context.Context, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *appsv1.DeploymentList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *appsv1.DeploymentList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.DeploymentList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *DeploymentInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appsv1.Deployment, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *appsv1.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *appsv1.Deployment); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, deployment, opts
func (_m *DeploymentInterface) Update(ctx context.Context, deployment *appsv1.Deployment, opts metav1.UpdateOptions) (*appsv1.Deployment, error) {
	ret := _m.Called(ctx, deployment, opts)

	var r0 *appsv1.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.Deployment, metav1.UpdateOptions) *appsv1.Deployment); ok {
		r0 = rf(ctx, deployment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.Deployment, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, deployment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateScale provides a mock function with given fields: ctx, deploymentName, scale, opts
func (_m *DeploymentInterface) UpdateScale(ctx context.Context, deploymentName string, scale *apiautoscalingv1.Scale, opts metav1.UpdateOptions) (*apiautoscalingv1.Scale, error) {
	ret := _m.Called(ctx, deploymentName, scale, opts)

	var r0 *apiautoscalingv1.Scale
	if rf, ok := ret.Get(0).(func(context.Context, string, *apiautoscalingv1.Scale, metav1.UpdateOptions) *apiautoscalingv1.Scale); ok {
		r0 = rf(ctx, deploymentName, scale, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiautoscalingv1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *apiautoscalingv1.Scale, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, deploymentName, scale, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, deployment, opts
func (_m *DeploymentInterface) UpdateStatus(ctx context.Context, deployment *appsv1.Deployment, opts metav1.UpdateOptions) (*appsv1.Deployment, error) {
	ret := _m.Called(ctx, deployment, opts)

	var r0 *appsv1.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.Deployment, metav1.UpdateOptions) *appsv1.Deployment); ok {
		r0 = rf(ctx, deployment, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.Deployment, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, deployment, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *DeploymentInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDeploymentInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeploymentInterface creates a new instance of DeploymentInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeploymentInterface(t mockConstructorTestingTNewDeploymentInterface) *DeploymentInterface {
	mock := &DeploymentInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
