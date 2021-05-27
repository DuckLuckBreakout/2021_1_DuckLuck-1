// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/favorites (interfaces: UseCase)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// AddProductToFavorites mocks base method.
func (m *MockUseCase) AddProductToFavorites(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProductToFavorites", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProductToFavorites indicates an expected call of AddProductToFavorites.
func (mr *MockUseCaseMockRecorder) AddProductToFavorites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProductToFavorites", reflect.TypeOf((*MockUseCase)(nil).AddProductToFavorites), arg0, arg1)
}

// DeleteProductFromFavorites mocks base method.
func (m *MockUseCase) DeleteProductFromFavorites(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductFromFavorites", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductFromFavorites indicates an expected call of DeleteProductFromFavorites.
func (mr *MockUseCaseMockRecorder) DeleteProductFromFavorites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductFromFavorites", reflect.TypeOf((*MockUseCase)(nil).DeleteProductFromFavorites), arg0, arg1)
}

// GetRangeFavorites mocks base method.
func (m *MockUseCase) GetRangeFavorites(arg0 *models.PaginatorFavorites, arg1 uint64) (*models.RangeFavorites, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRangeFavorites", arg0, arg1)
	ret0, _ := ret[0].(*models.RangeFavorites)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRangeFavorites indicates an expected call of GetRangeFavorites.
func (mr *MockUseCaseMockRecorder) GetRangeFavorites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRangeFavorites", reflect.TypeOf((*MockUseCase)(nil).GetRangeFavorites), arg0, arg1)
}

// GetUserFavorites mocks base method.
func (m *MockUseCase) GetUserFavorites(arg0 uint64) (*models.UserFavorites, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFavorites", arg0)
	ret0, _ := ret[0].(*models.UserFavorites)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFavorites indicates an expected call of GetUserFavorites.
func (mr *MockUseCaseMockRecorder) GetUserFavorites(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFavorites", reflect.TypeOf((*MockUseCase)(nil).GetUserFavorites), arg0)
}
