// Code generated by MockGen. DO NOT EDIT.
// Source: internal/codegen/generator.go

// Package codegen_test is a generated GoMock package.
package codegen_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	afero "github.com/spf13/afero"
)

// MockAddGenerator is a mock of AddGenerator interface.
type MockAddGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockAddGeneratorMockRecorder
}

// MockAddGeneratorMockRecorder is the mock recorder for MockAddGenerator.
type MockAddGeneratorMockRecorder struct {
	mock *MockAddGenerator
}

// NewMockAddGenerator creates a new mock instance.
func NewMockAddGenerator(ctrl *gomock.Controller) *MockAddGenerator {
	mock := &MockAddGenerator{ctrl: ctrl}
	mock.recorder = &MockAddGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddGenerator) EXPECT() *MockAddGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockAddGenerator) Generate(arg0 afero.Fs, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockAddGeneratorMockRecorder) Generate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockAddGenerator)(nil).Generate), arg0, arg1, arg2)
}
