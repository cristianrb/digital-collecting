// Code generated by MockGen. DO NOT EDIT.
// Source: dc-backend/internal/token (interfaces: JWTValidator)

// Package mocktoken is a generated GoMock package.
package mocktoken

import (
	token "dc-backend/internal/token"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJWTValidator is a mock of JWTValidator interface.
type MockJWTValidator struct {
	ctrl     *gomock.Controller
	recorder *MockJWTValidatorMockRecorder
}

// MockJWTValidatorMockRecorder is the mock recorder for MockJWTValidator.
type MockJWTValidatorMockRecorder struct {
	mock *MockJWTValidator
}

// NewMockJWTValidator creates a new mock instance.
func NewMockJWTValidator(ctrl *gomock.Controller) *MockJWTValidator {
	mock := &MockJWTValidator{ctrl: ctrl}
	mock.recorder = &MockJWTValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTValidator) EXPECT() *MockJWTValidatorMockRecorder {
	return m.recorder
}

// VerifyToken mocks base method.
func (m *MockJWTValidator) VerifyToken(arg0 string) (*token.Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", arg0)
	ret0, _ := ret[0].(*token.Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockJWTValidatorMockRecorder) VerifyToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockJWTValidator)(nil).VerifyToken), arg0)
}