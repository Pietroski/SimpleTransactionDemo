// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/devices/sqlc (interfaces: Store)

// Package mockedDeviceStore is a generated GoMock package.
package mockedDeviceStore

import (
	context "context"
	reflect "reflect"

	sqlc_device_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/devices/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockStore) CreateDevice(arg0 context.Context, arg1 sqlc_device_store.CreateDeviceParams) (sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", arg0, arg1)
	ret0, _ := ret[0].(sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockStoreMockRecorder) CreateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockStore)(nil).CreateDevice), arg0, arg1)
}

// DeleteDevice mocks base method.
func (m *MockStore) DeleteDevice(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockStoreMockRecorder) DeleteDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockStore)(nil).DeleteDevice), arg0, arg1)
}

// DeleteUserDevices mocks base method.
func (m *MockStore) DeleteUserDevices(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserDevices", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserDevices indicates an expected call of DeleteUserDevices.
func (mr *MockStoreMockRecorder) DeleteUserDevices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserDevices", reflect.TypeOf((*MockStore)(nil).DeleteUserDevices), arg0, arg1)
}

// GetDevice mocks base method.
func (m *MockStore) GetDevice(arg0 context.Context, arg1 uuid.UUID) (sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevice", arg0, arg1)
	ret0, _ := ret[0].(sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice.
func (mr *MockStoreMockRecorder) GetDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockStore)(nil).GetDevice), arg0, arg1)
}

// GetPaginatedUserDevices mocks base method.
func (m *MockStore) GetPaginatedUserDevices(arg0 context.Context, arg1 sqlc_device_store.GetPaginatedUserDevicesParams) (sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaginatedUserDevices", arg0, arg1)
	ret0, _ := ret[0].(sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaginatedUserDevices indicates an expected call of GetPaginatedUserDevices.
func (mr *MockStoreMockRecorder) GetPaginatedUserDevices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaginatedUserDevices", reflect.TypeOf((*MockStore)(nil).GetPaginatedUserDevices), arg0, arg1)
}

// GetUserDevices mocks base method.
func (m *MockStore) GetUserDevices(arg0 context.Context, arg1 uuid.UUID) (sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDevices", arg0, arg1)
	ret0, _ := ret[0].(sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDevices indicates an expected call of GetUserDevices.
func (mr *MockStoreMockRecorder) GetUserDevices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDevices", reflect.TypeOf((*MockStore)(nil).GetUserDevices), arg0, arg1)
}

// ListDevices mocks base method.
func (m *MockStore) ListDevices(arg0 context.Context) ([]sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDevices", arg0)
	ret0, _ := ret[0].([]sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDevices indicates an expected call of ListDevices.
func (mr *MockStoreMockRecorder) ListDevices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDevices", reflect.TypeOf((*MockStore)(nil).ListDevices), arg0)
}

// ListPaginatedDevices mocks base method.
func (m *MockStore) ListPaginatedDevices(arg0 context.Context, arg1 sqlc_device_store.ListPaginatedDevicesParams) ([]sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPaginatedDevices", arg0, arg1)
	ret0, _ := ret[0].([]sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPaginatedDevices indicates an expected call of ListPaginatedDevices.
func (mr *MockStoreMockRecorder) ListPaginatedDevices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPaginatedDevices", reflect.TypeOf((*MockStore)(nil).ListPaginatedDevices), arg0, arg1)
}

// UpdateDevice mocks base method.
func (m *MockStore) UpdateDevice(arg0 context.Context, arg1 sqlc_device_store.UpdateDeviceParams) (sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDevice", arg0, arg1)
	ret0, _ := ret[0].(sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockStoreMockRecorder) UpdateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockStore)(nil).UpdateDevice), arg0, arg1)
}

// UpdateUserDevice mocks base method.
func (m *MockStore) UpdateUserDevice(arg0 context.Context, arg1 sqlc_device_store.UpdateUserDeviceParams) (sqlc_device_store.Devices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDevice", arg0, arg1)
	ret0, _ := ret[0].(sqlc_device_store.Devices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserDevice indicates an expected call of UpdateUserDevice.
func (mr *MockStoreMockRecorder) UpdateUserDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDevice", reflect.TypeOf((*MockStore)(nil).UpdateUserDevice), arg0, arg1)
}