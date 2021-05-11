// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/product (interfaces: Repository)

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

// CreateFilterString mocks base method.
func (m *MockRepository) CreateFilterString(arg0 *models.ProductFilter) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFilterString", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateFilterString indicates an expected call of CreateFilterString.
func (mr *MockRepositoryMockRecorder) CreateFilterString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilterString", reflect.TypeOf((*MockRepository)(nil).CreateFilterString), arg0)
}

// CreateSortString mocks base method.
func (m *MockRepository) CreateSortString(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSortString", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSortString indicates an expected call of CreateSortString.
func (mr *MockRepositoryMockRecorder) CreateSortString(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSortString", reflect.TypeOf((*MockRepository)(nil).CreateSortString), arg0, arg1)
}

// GetCountPages mocks base method.
func (m *MockRepository) GetCountPages(arg0 uint64, arg1 int, arg2 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountPages", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountPages indicates an expected call of GetCountPages.
func (mr *MockRepositoryMockRecorder) GetCountPages(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountPages", reflect.TypeOf((*MockRepository)(nil).GetCountPages), arg0, arg1, arg2)
}

// GetCountSearchPages mocks base method.
func (m *MockRepository) GetCountSearchPages(arg0 uint64, arg1 int, arg2, arg3 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountSearchPages", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountSearchPages indicates an expected call of GetCountSearchPages.
func (mr *MockRepositoryMockRecorder) GetCountSearchPages(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountSearchPages", reflect.TypeOf((*MockRepository)(nil).GetCountSearchPages), arg0, arg1, arg2, arg3)
}

// SearchRangeProducts mocks base method.
func (m *MockRepository) SearchRangeProducts(arg0 *models.SearchQuery, arg1, arg2 string) ([]*models.ViewProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRangeProducts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.ViewProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRangeProducts indicates an expected call of SearchRangeProducts.
func (mr *MockRepositoryMockRecorder) SearchRangeProducts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRangeProducts", reflect.TypeOf((*MockRepository)(nil).SearchRangeProducts), arg0, arg1, arg2)
}

// SelectProductById mocks base method.
func (m *MockRepository) SelectProductById(arg0 uint64) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectProductById", arg0)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectProductById indicates an expected call of SelectProductById.
func (mr *MockRepositoryMockRecorder) SelectProductById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectProductById", reflect.TypeOf((*MockRepository)(nil).SelectProductById), arg0)
}

// SelectRangeProducts mocks base method.
func (m *MockRepository) SelectRangeProducts(arg0 *models.PaginatorProducts, arg1, arg2 string) ([]*models.ViewProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectRangeProducts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.ViewProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectRangeProducts indicates an expected call of SelectRangeProducts.
func (mr *MockRepositoryMockRecorder) SelectRangeProducts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectRangeProducts", reflect.TypeOf((*MockRepository)(nil).SelectRangeProducts), arg0, arg1, arg2)
}

// SelectRecommendationsByCategory mocks base method.
func (m *MockRepository) SelectRecommendationsByCategory(arg0 uint64, arg1 int) ([]*models.RecommendationProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectRecommendationsByCategory", arg0, arg1)
	ret0, _ := ret[0].([]*models.RecommendationProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectRecommendationsByCategory indicates an expected call of SelectRecommendationsByCategory.
func (mr *MockRepositoryMockRecorder) SelectRecommendationsByCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectRecommendationsByCategory", reflect.TypeOf((*MockRepository)(nil).SelectRecommendationsByCategory), arg0, arg1)
}

// SelectRecommendationsByReviews mocks base method.
func (m *MockRepository) SelectRecommendationsByReviews(arg0 uint64, arg1 int) ([]*models.RecommendationProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectRecommendationsByReviews", arg0, arg1)
	ret0, _ := ret[0].([]*models.RecommendationProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectRecommendationsByReviews indicates an expected call of SelectRecommendationsByReviews.
func (mr *MockRepositoryMockRecorder) SelectRecommendationsByReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectRecommendationsByReviews", reflect.TypeOf((*MockRepository)(nil).SelectRecommendationsByReviews), arg0, arg1)
}
