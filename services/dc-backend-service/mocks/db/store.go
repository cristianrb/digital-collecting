// Code generated by MockGen. DO NOT EDIT.
// Source: dc-backend/internal/storage (interfaces: ItemStorage)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	types "dc-backend/pkg/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockItemStorage is a mock of ItemStorage interface.
type MockItemStorage struct {
	ctrl     *gomock.Controller
	recorder *MockItemStorageMockRecorder
}

// MockItemStorageMockRecorder is the mock recorder for MockItemStorage.
type MockItemStorageMockRecorder struct {
	mock *MockItemStorage
}

// NewMockItemStorage creates a new mock instance.
func NewMockItemStorage(ctrl *gomock.Controller) *MockItemStorage {
	mock := &MockItemStorage{ctrl: ctrl}
	mock.recorder = &MockItemStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemStorage) EXPECT() *MockItemStorageMockRecorder {
	return m.recorder
}

// GetAllItems mocks base method.
func (m *MockItemStorage) GetAllItems(arg0, arg1 int) (types.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllItems", arg0, arg1)
	ret0, _ := ret[0].(types.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllItems indicates an expected call of GetAllItems.
func (mr *MockItemStorageMockRecorder) GetAllItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllItems", reflect.TypeOf((*MockItemStorage)(nil).GetAllItems), arg0, arg1)
}
