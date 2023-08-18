// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	sonar "github.com/epam/edp-sonar-operator/pkg/client/sonar"
	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterface struct {
	mock.Mock
}

// AddGroupToPermissionTemplate provides a mock function with given fields: ctx, templateID, permGroup
func (_m *ClientInterface) AddGroupToPermissionTemplate(ctx context.Context, templateID string, permGroup *sonar.PermissionTemplateGroup) error {
	ret := _m.Called(ctx, templateID, permGroup)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *sonar.PermissionTemplateGroup) error); ok {
		r0 = rf(ctx, templateID, permGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPermissionToUser provides a mock function with given fields: ctx, userLogin, permission
func (_m *ClientInterface) AddPermissionToUser(ctx context.Context, userLogin string, permission string) error {
	ret := _m.Called(ctx, userLogin, permission)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userLogin, permission)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPermissionsToGroup provides a mock function with given fields: groupName, permissions
func (_m *ClientInterface) AddPermissionsToGroup(groupName string, permissions string) error {
	ret := _m.Called(groupName, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(groupName, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUserToGroup provides a mock function with given fields: ctx, userLogin, groupName
func (_m *ClientInterface) AddUserToGroup(ctx context.Context, userLogin string, groupName string) error {
	ret := _m.Called(ctx, userLogin, groupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userLogin, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfigureGeneralSettings provides a mock function with given fields: settings
func (_m *ClientInterface) ConfigureGeneralSettings(settings ...sonar.SettingRequest) error {
	_va := make([]interface{}, len(settings))
	for _i := range settings {
		_va[_i] = settings[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...sonar.SettingRequest) error); ok {
		r0 = rf(settings...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateGroup provides a mock function with given fields: ctx, gr
func (_m *ClientInterface) CreateGroup(ctx context.Context, gr *sonar.Group) error {
	ret := _m.Called(ctx, gr)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sonar.Group) error); ok {
		r0 = rf(ctx, gr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePermissionTemplate provides a mock function with given fields: ctx, tpl
func (_m *ClientInterface) CreatePermissionTemplate(ctx context.Context, tpl *sonar.PermissionTemplateData) (string, error) {
	ret := _m.Called(ctx, tpl)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sonar.PermissionTemplateData) (string, error)); ok {
		return rf(ctx, tpl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sonar.PermissionTemplateData) string); ok {
		r0 = rf(ctx, tpl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sonar.PermissionTemplateData) error); ok {
		r1 = rf(ctx, tpl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQualityGate provides a mock function with given fields: ctx, name
func (_m *ClientInterface) CreateQualityGate(ctx context.Context, name string) (*sonar.QualityGate, error) {
	ret := _m.Called(ctx, name)

	var r0 *sonar.QualityGate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sonar.QualityGate, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sonar.QualityGate); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sonar.QualityGate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQualityGateCondition provides a mock function with given fields: ctx, gate, condition
func (_m *ClientInterface) CreateQualityGateCondition(ctx context.Context, gate string, condition sonar.QualityGateCondition) error {
	ret := _m.Called(ctx, gate, condition)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, sonar.QualityGateCondition) error); ok {
		r0 = rf(ctx, gate, condition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *ClientInterface) CreateUser(ctx context.Context, u *sonar.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sonar.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeactivateUser provides a mock function with given fields: ctx, userLogin
func (_m *ClientInterface) DeactivateUser(ctx context.Context, userLogin string) error {
	ret := _m.Called(ctx, userLogin)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userLogin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGroup provides a mock function with given fields: ctx, groupName
func (_m *ClientInterface) DeleteGroup(ctx context.Context, groupName string) error {
	ret := _m.Called(ctx, groupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePermissionTemplate provides a mock function with given fields: ctx, id
func (_m *ClientInterface) DeletePermissionTemplate(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteQualityGate provides a mock function with given fields: ctx, name
func (_m *ClientInterface) DeleteQualityGate(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteQualityGateCondition provides a mock function with given fields: ctx, conditionId
func (_m *ClientInterface) DeleteQualityGateCondition(ctx context.Context, conditionId string) error {
	ret := _m.Called(ctx, conditionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, conditionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateUserToken provides a mock function with given fields: userName
func (_m *ClientInterface) GenerateUserToken(userName string) (*string, error) {
	ret := _m.Called(userName)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*string, error)); ok {
		return rf(userName)
	}
	if rf, ok := ret.Get(0).(func(string) *string); ok {
		r0 = rf(userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroup provides a mock function with given fields: ctx, groupName
func (_m *ClientInterface) GetGroup(ctx context.Context, groupName string) (*sonar.Group, error) {
	ret := _m.Called(ctx, groupName)

	var r0 *sonar.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sonar.Group, error)); ok {
		return rf(ctx, groupName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sonar.Group); ok {
		r0 = rf(ctx, groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sonar.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, groupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPermissionTemplate provides a mock function with given fields: ctx, name
func (_m *ClientInterface) GetPermissionTemplate(ctx context.Context, name string) (*sonar.PermissionTemplate, error) {
	ret := _m.Called(ctx, name)

	var r0 *sonar.PermissionTemplate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sonar.PermissionTemplate, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sonar.PermissionTemplate); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sonar.PermissionTemplate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPermissionTemplateGroups provides a mock function with given fields: ctx, templateID
func (_m *ClientInterface) GetPermissionTemplateGroups(ctx context.Context, templateID string) ([]sonar.PermissionTemplateGroup, error) {
	ret := _m.Called(ctx, templateID)

	var r0 []sonar.PermissionTemplateGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]sonar.PermissionTemplateGroup, error)); ok {
		return rf(ctx, templateID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []sonar.PermissionTemplateGroup); ok {
		r0 = rf(ctx, templateID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sonar.PermissionTemplateGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, templateID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQualityGate provides a mock function with given fields: ctx, name
func (_m *ClientInterface) GetQualityGate(ctx context.Context, name string) (*sonar.QualityGate, error) {
	ret := _m.Called(ctx, name)

	var r0 *sonar.QualityGate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sonar.QualityGate, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sonar.QualityGate); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sonar.QualityGate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByLogin provides a mock function with given fields: ctx, userLogin
func (_m *ClientInterface) GetUserByLogin(ctx context.Context, userLogin string) (*sonar.User, error) {
	ret := _m.Called(ctx, userLogin)

	var r0 *sonar.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sonar.User, error)); ok {
		return rf(ctx, userLogin)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sonar.User); ok {
		r0 = rf(ctx, userLogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sonar.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userLogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserGroups provides a mock function with given fields: ctx, userLogin
func (_m *ClientInterface) GetUserGroups(ctx context.Context, userLogin string) ([]sonar.Group, error) {
	ret := _m.Called(ctx, userLogin)

	var r0 []sonar.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]sonar.Group, error)); ok {
		return rf(ctx, userLogin)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []sonar.Group); ok {
		r0 = rf(ctx, userLogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sonar.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userLogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserPermissions provides a mock function with given fields: ctx, userLogin
func (_m *ClientInterface) GetUserPermissions(ctx context.Context, userLogin string) ([]string, error) {
	ret := _m.Called(ctx, userLogin)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, userLogin)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, userLogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userLogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserToken provides a mock function with given fields: ctx, userLogin, tokenName
func (_m *ClientInterface) GetUserToken(ctx context.Context, userLogin string, tokenName string) (*sonar.UserToken, error) {
	ret := _m.Called(ctx, userLogin, tokenName)

	var r0 *sonar.UserToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*sonar.UserToken, error)); ok {
		return rf(ctx, userLogin, tokenName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *sonar.UserToken); ok {
		r0 = rf(ctx, userLogin, tokenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sonar.UserToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userLogin, tokenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: ctx
func (_m *ClientInterface) Health(ctx context.Context) (*sonar.SystemHealth, error) {
	ret := _m.Called(ctx)

	var r0 *sonar.SystemHealth
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*sonar.SystemHealth, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *sonar.SystemHealth); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sonar.SystemHealth)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallPlugins provides a mock function with given fields: plugins
func (_m *ClientInterface) InstallPlugins(plugins []string) error {
	ret := _m.Called(plugins)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(plugins)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveGroupFromPermissionTemplate provides a mock function with given fields: ctx, templateID, permGroup
func (_m *ClientInterface) RemoveGroupFromPermissionTemplate(ctx context.Context, templateID string, permGroup *sonar.PermissionTemplateGroup) error {
	ret := _m.Called(ctx, templateID, permGroup)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *sonar.PermissionTemplateGroup) error); ok {
		r0 = rf(ctx, templateID, permGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePermissionFromUser provides a mock function with given fields: ctx, userLogin, permission
func (_m *ClientInterface) RemovePermissionFromUser(ctx context.Context, userLogin string, permission string) error {
	ret := _m.Called(ctx, userLogin, permission)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userLogin, permission)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUserFromGroup provides a mock function with given fields: ctx, userLogin, groupName
func (_m *ClientInterface) RemoveUserFromGroup(ctx context.Context, userLogin string, groupName string) error {
	ret := _m.Called(ctx, userLogin, groupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userLogin, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetSettings provides a mock function with given fields: ctx, settingsKeys
func (_m *ClientInterface) ResetSettings(ctx context.Context, settingsKeys []string) error {
	ret := _m.Called(ctx, settingsKeys)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, settingsKeys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAsDefaultQualityGate provides a mock function with given fields: ctx, name
func (_m *ClientInterface) SetAsDefaultQualityGate(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDefaultPermissionTemplate provides a mock function with given fields: ctx, name
func (_m *ClientInterface) SetDefaultPermissionTemplate(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetProjectsDefaultVisibility provides a mock function with given fields: visibility
func (_m *ClientInterface) SetProjectsDefaultVisibility(visibility string) error {
	ret := _m.Called(visibility)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(visibility)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSetting provides a mock function with given fields: ctx, setting
func (_m *ClientInterface) SetSetting(ctx context.Context, setting url.Values) error {
	ret := _m.Called(ctx, setting)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, url.Values) error); ok {
		r0 = rf(ctx, setting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateGroup provides a mock function with given fields: ctx, currentName, group
func (_m *ClientInterface) UpdateGroup(ctx context.Context, currentName string, group *sonar.Group) error {
	ret := _m.Called(ctx, currentName, group)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *sonar.Group) error); ok {
		r0 = rf(ctx, currentName, group)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePermissionTemplate provides a mock function with given fields: ctx, tpl
func (_m *ClientInterface) UpdatePermissionTemplate(ctx context.Context, tpl *sonar.PermissionTemplate) error {
	ret := _m.Called(ctx, tpl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sonar.PermissionTemplate) error); ok {
		r0 = rf(ctx, tpl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateQualityGateCondition provides a mock function with given fields: ctx, condition
func (_m *ClientInterface) UpdateQualityGateCondition(ctx context.Context, condition sonar.QualityGateCondition) error {
	ret := _m.Called(ctx, condition)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sonar.QualityGateCondition) error); ok {
		r0 = rf(ctx, condition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, u
func (_m *ClientInterface) UpdateUser(ctx context.Context, u *sonar.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sonar.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClientInterface creates a new instance of ClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientInterface {
	mock := &ClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
