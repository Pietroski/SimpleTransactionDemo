// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/transactions (interfaces: Store)

// Package mockedBankAccountStore is a generated GoMock package.
package mockedBankAccountStore

import (
	context "context"
	reflect "reflect"

	bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/bank-accounts"
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

// GetPaginatedUserWallets mocks base method.
func (m *MockStore) GetPaginatedUserWallets(arg0 context.Context, arg1 bank_account_store.GetPaginatedUserWalletsParams) ([]bank_account_store.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaginatedUserWallets", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaginatedUserWallets indicates an expected call of GetPaginatedUserWallets.
func (mr *MockStoreMockRecorder) GetPaginatedUserWallets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaginatedUserWallets", reflect.TypeOf((*MockStore)(nil).GetPaginatedUserWallets), arg0, arg1)
}

// GetPaginatedWalletsByUserID mocks base method.
func (m *MockStore) GetPaginatedWalletsByUserID(arg0 context.Context, arg1 bank_account_store.GetPaginatedWalletsByUserIDParams) ([]bank_account_store.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaginatedWalletsByUserID", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaginatedWalletsByUserID indicates an expected call of GetPaginatedWalletsByUserID.
func (mr *MockStoreMockRecorder) GetPaginatedWalletsByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaginatedWalletsByUserID", reflect.TypeOf((*MockStore)(nil).GetPaginatedWalletsByUserID), arg0, arg1)
}

// GetUserWallets mocks base method.
func (m *MockStore) GetUserWallets(arg0 context.Context, arg1 uuid.UUID) ([]bank_account_store.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWallets", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWallets indicates an expected call of GetUserWallets.
func (mr *MockStoreMockRecorder) GetUserWallets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWallets", reflect.TypeOf((*MockStore)(nil).GetUserWallets), arg0, arg1)
}

// GetWalletsByUserID mocks base method.
func (m *MockStore) GetWalletsByUserID(arg0 context.Context, arg1 uuid.UUID) ([]bank_account_store.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletsByUserID", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletsByUserID indicates an expected call of GetWalletsByUserID.
func (mr *MockStoreMockRecorder) GetWalletsByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletsByUserID", reflect.TypeOf((*MockStore)(nil).GetWalletsByUserID), arg0, arg1)
}

// ListFromUserTransactionLogs mocks base method.
func (m *MockStore) ListFromUserTransactionLogs(arg0 context.Context, arg1 uuid.UUID) ([]bank_account_store.TransactionRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFromUserTransactionLogs", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.TransactionRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFromUserTransactionLogs indicates an expected call of ListFromUserTransactionLogs.
func (mr *MockStoreMockRecorder) ListFromUserTransactionLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFromUserTransactionLogs", reflect.TypeOf((*MockStore)(nil).ListFromUserTransactionLogs), arg0, arg1)
}

// ListPaginatedFromUserTransactionLogs mocks base method.
func (m *MockStore) ListPaginatedFromUserTransactionLogs(arg0 context.Context, arg1 bank_account_store.ListPaginatedFromUserTransactionLogsParams) ([]bank_account_store.TransactionRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPaginatedFromUserTransactionLogs", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.TransactionRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPaginatedFromUserTransactionLogs indicates an expected call of ListPaginatedFromUserTransactionLogs.
func (mr *MockStoreMockRecorder) ListPaginatedFromUserTransactionLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPaginatedFromUserTransactionLogs", reflect.TypeOf((*MockStore)(nil).ListPaginatedFromUserTransactionLogs), arg0, arg1)
}

// ListPaginatedToUserTransactionLogs mocks base method.
func (m *MockStore) ListPaginatedToUserTransactionLogs(arg0 context.Context, arg1 bank_account_store.ListPaginatedToUserTransactionLogsParams) ([]bank_account_store.TransactionRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPaginatedToUserTransactionLogs", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.TransactionRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPaginatedToUserTransactionLogs indicates an expected call of ListPaginatedToUserTransactionLogs.
func (mr *MockStoreMockRecorder) ListPaginatedToUserTransactionLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPaginatedToUserTransactionLogs", reflect.TypeOf((*MockStore)(nil).ListPaginatedToUserTransactionLogs), arg0, arg1)
}

// ListPaginatedTransactionLogs mocks base method.
func (m *MockStore) ListPaginatedTransactionLogs(arg0 context.Context, arg1 bank_account_store.ListPaginatedTransactionLogsParams) ([]bank_account_store.TransactionRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPaginatedTransactionLogs", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.TransactionRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPaginatedTransactionLogs indicates an expected call of ListPaginatedTransactionLogs.
func (mr *MockStoreMockRecorder) ListPaginatedTransactionLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPaginatedTransactionLogs", reflect.TypeOf((*MockStore)(nil).ListPaginatedTransactionLogs), arg0, arg1)
}

// ListToUserTransactionLogs mocks base method.
func (m *MockStore) ListToUserTransactionLogs(arg0 context.Context, arg1 uuid.UUID) ([]bank_account_store.TransactionRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListToUserTransactionLogs", arg0, arg1)
	ret0, _ := ret[0].([]bank_account_store.TransactionRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListToUserTransactionLogs indicates an expected call of ListToUserTransactionLogs.
func (mr *MockStoreMockRecorder) ListToUserTransactionLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListToUserTransactionLogs", reflect.TypeOf((*MockStore)(nil).ListToUserTransactionLogs), arg0, arg1)
}

// ListTransactionLogs mocks base method.
func (m *MockStore) ListTransactionLogs(arg0 context.Context) ([]bank_account_store.TransactionRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransactionLogs", arg0)
	ret0, _ := ret[0].([]bank_account_store.TransactionRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransactionLogs indicates an expected call of ListTransactionLogs.
func (mr *MockStoreMockRecorder) ListTransactionLogs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransactionLogs", reflect.TypeOf((*MockStore)(nil).ListTransactionLogs), arg0)
}

// LogTransaction mocks base method.
func (m *MockStore) LogTransaction(arg0 context.Context, arg1 bank_account_store.LogTransactionParams) (bank_account_store.TransactionRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogTransaction", arg0, arg1)
	ret0, _ := ret[0].(bank_account_store.TransactionRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogTransaction indicates an expected call of LogTransaction.
func (mr *MockStoreMockRecorder) LogTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogTransaction", reflect.TypeOf((*MockStore)(nil).LogTransaction), arg0, arg1)
}

// UpdateUserWalletBalance mocks base method.
func (m *MockStore) UpdateUserWalletBalance(arg0 context.Context, arg1 bank_account_store.UpdateUserWalletBalanceParams) (bank_account_store.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserWalletBalance", arg0, arg1)
	ret0, _ := ret[0].(bank_account_store.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserWalletBalance indicates an expected call of UpdateUserWalletBalance.
func (mr *MockStoreMockRecorder) UpdateUserWalletBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserWalletBalance", reflect.TypeOf((*MockStore)(nil).UpdateUserWalletBalance), arg0, arg1)
}