// Code generated by MockGen. DO NOT EDIT.
// Source: customer.go

// Package mock_customer is a generated GoMock package.
package mock_customer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomer is a mock of Customer interface.
type MockCustomer struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerMockRecorder
}

// MockCustomerMockRecorder is the mock recorder for MockCustomer.
type MockCustomerMockRecorder struct {
	mock *MockCustomer
}

// NewMockCustomer creates a new mock instance.
func NewMockCustomer(ctrl *gomock.Controller) *MockCustomer {
	mock := &MockCustomer{ctrl: ctrl}
	mock.recorder = &MockCustomerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomer) EXPECT() *MockCustomerMockRecorder {
	return m.recorder
}

// GetCustomerId mocks base method.
func (m *MockCustomer) GetCustomerId() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerId")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCustomerId indicates an expected call of GetCustomerId.
func (mr *MockCustomerMockRecorder) GetCustomerId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerId", reflect.TypeOf((*MockCustomer)(nil).GetCustomerId))
}

// HaircutCompleted mocks base method.
func (m *MockCustomer) HaircutCompleted() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HaircutCompleted")
}

// HaircutCompleted indicates an expected call of HaircutCompleted.
func (mr *MockCustomerMockRecorder) HaircutCompleted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HaircutCompleted", reflect.TypeOf((*MockCustomer)(nil).HaircutCompleted))
}

// WaitForHaircutToBeCompleted mocks base method.
func (m *MockCustomer) WaitForHaircutToBeCompleted() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WaitForHaircutToBeCompleted")
}

// WaitForHaircutToBeCompleted indicates an expected call of WaitForHaircutToBeCompleted.
func (mr *MockCustomerMockRecorder) WaitForHaircutToBeCompleted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForHaircutToBeCompleted", reflect.TypeOf((*MockCustomer)(nil).WaitForHaircutToBeCompleted))
}