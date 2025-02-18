// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChainSafe/gossamer/dot/rpc (interfaces: API,TransactionStateAPI)

// Package rpc is a generated GoMock package.
package rpc

import (
	reflect "reflect"

	types "github.com/ChainSafe/gossamer/dot/types"
	common "github.com/ChainSafe/gossamer/lib/common"
	transaction "github.com/ChainSafe/gossamer/lib/transaction"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// BuildMethodNames mocks base method.
func (m *MockAPI) BuildMethodNames(arg0 interface{}, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BuildMethodNames", arg0, arg1)
}

// BuildMethodNames indicates an expected call of BuildMethodNames.
func (mr *MockAPIMockRecorder) BuildMethodNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildMethodNames", reflect.TypeOf((*MockAPI)(nil).BuildMethodNames), arg0, arg1)
}

// Methods mocks base method.
func (m *MockAPI) Methods() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Methods")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Methods indicates an expected call of Methods.
func (mr *MockAPIMockRecorder) Methods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Methods", reflect.TypeOf((*MockAPI)(nil).Methods))
}

// MockTransactionStateAPI is a mock of TransactionStateAPI interface.
type MockTransactionStateAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionStateAPIMockRecorder
}

// MockTransactionStateAPIMockRecorder is the mock recorder for MockTransactionStateAPI.
type MockTransactionStateAPIMockRecorder struct {
	mock *MockTransactionStateAPI
}

// NewMockTransactionStateAPI creates a new mock instance.
func NewMockTransactionStateAPI(ctrl *gomock.Controller) *MockTransactionStateAPI {
	mock := &MockTransactionStateAPI{ctrl: ctrl}
	mock.recorder = &MockTransactionStateAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionStateAPI) EXPECT() *MockTransactionStateAPIMockRecorder {
	return m.recorder
}

// AddToPool mocks base method.
func (m *MockTransactionStateAPI) AddToPool(arg0 *transaction.ValidTransaction) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToPool", arg0)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// AddToPool indicates an expected call of AddToPool.
func (mr *MockTransactionStateAPIMockRecorder) AddToPool(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToPool", reflect.TypeOf((*MockTransactionStateAPI)(nil).AddToPool), arg0)
}

// FreeStatusNotifierChannel mocks base method.
func (m *MockTransactionStateAPI) FreeStatusNotifierChannel(arg0 chan transaction.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FreeStatusNotifierChannel", arg0)
}

// FreeStatusNotifierChannel indicates an expected call of FreeStatusNotifierChannel.
func (mr *MockTransactionStateAPIMockRecorder) FreeStatusNotifierChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeStatusNotifierChannel", reflect.TypeOf((*MockTransactionStateAPI)(nil).FreeStatusNotifierChannel), arg0)
}

// GetStatusNotifierChannel mocks base method.
func (m *MockTransactionStateAPI) GetStatusNotifierChannel(arg0 types.Extrinsic) chan transaction.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatusNotifierChannel", arg0)
	ret0, _ := ret[0].(chan transaction.Status)
	return ret0
}

// GetStatusNotifierChannel indicates an expected call of GetStatusNotifierChannel.
func (mr *MockTransactionStateAPIMockRecorder) GetStatusNotifierChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatusNotifierChannel", reflect.TypeOf((*MockTransactionStateAPI)(nil).GetStatusNotifierChannel), arg0)
}

// Pending mocks base method.
func (m *MockTransactionStateAPI) Pending() []*transaction.ValidTransaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pending")
	ret0, _ := ret[0].([]*transaction.ValidTransaction)
	return ret0
}

// Pending indicates an expected call of Pending.
func (mr *MockTransactionStateAPIMockRecorder) Pending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pending", reflect.TypeOf((*MockTransactionStateAPI)(nil).Pending))
}
