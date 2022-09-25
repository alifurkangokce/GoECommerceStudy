// Code generated by MockGen. DO NOT EDIT.
// Source: GoECommerceStudy/services (interfaces: ProductService)

// Package services is a generated GoMock package.
package services

import (
	dto "GoECommerceStudy/dto"
	models "GoECommerceStudy/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// ProductDelete mocks base method.
func (m *MockProductService) ProductDelete(arg0 primitive.ObjectID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProductDelete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductDelete indicates an expected call of ProductDelete.
func (mr *MockProductServiceMockRecorder) ProductDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductDelete", reflect.TypeOf((*MockProductService)(nil).ProductDelete), arg0)
}

// ProductInsert mocks base method.
func (m *MockProductService) ProductInsert(arg0 models.Product) (*dto.ProductDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProductInsert", arg0)
	ret0, _ := ret[0].(*dto.ProductDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductInsert indicates an expected call of ProductInsert.
func (mr *MockProductServiceMockRecorder) ProductInsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductInsert", reflect.TypeOf((*MockProductService)(nil).ProductInsert), arg0)
}

// ProductsGet mocks base method.
func (m *MockProductService) ProductsGet() ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProductsGet")
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductsGet indicates an expected call of ProductsGet.
func (mr *MockProductServiceMockRecorder) ProductsGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductsGet", reflect.TypeOf((*MockProductService)(nil).ProductsGet))
}
