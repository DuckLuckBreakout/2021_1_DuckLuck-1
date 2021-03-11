// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/session (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRepository) Add(arg0 *models.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRepositoryMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepository)(nil).Add), arg0)
}

// DestroyByValue mocks base method.
func (m *MockRepository) DestroyByValue(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyByValue", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyByValue indicates an expected call of DestroyByValue.
func (mr *MockRepositoryMockRecorder) DestroyByValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyByValue", reflect.TypeOf((*MockRepository)(nil).DestroyByValue), arg0)
}

// GetByValue mocks base method.
func (m *MockRepository) GetByValue(arg0 string) (*models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByValue", arg0)
	ret0, _ := ret[0].(*models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByValue indicates an expected call of GetByValue.
func (mr *MockRepositoryMockRecorder) GetByValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByValue", reflect.TypeOf((*MockRepository)(nil).GetByValue), arg0)
}
