// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cycloidio/terracost/product (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	product "github.com/cycloidio/terracost/product"
	gomock "github.com/golang/mock/gomock"
)

// ProductRepository is a mock of Repository interface
type ProductRepository struct {
	ctrl     *gomock.Controller
	recorder *ProductRepositoryMockRecorder
}

// ProductRepositoryMockRecorder is the mock recorder for ProductRepository
type ProductRepositoryMockRecorder struct {
	mock *ProductRepository
}

// NewProductRepository creates a new mock instance
func NewProductRepository(ctrl *gomock.Controller) *ProductRepository {
	mock := &ProductRepository{ctrl: ctrl}
	mock.recorder = &ProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *ProductRepository) EXPECT() *ProductRepositoryMockRecorder {
	return m.recorder
}

// Filter mocks base method
func (m *ProductRepository) Filter(arg0 context.Context, arg1 *product.Filter) ([]*product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1)
	ret0, _ := ret[0].([]*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *ProductRepositoryMockRecorder) Filter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*ProductRepository)(nil).Filter), arg0, arg1)
}

// FindByVendorAndSKU mocks base method
func (m *ProductRepository) FindByVendorAndSKU(arg0 context.Context, arg1, arg2 string) (*product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByVendorAndSKU", arg0, arg1, arg2)
	ret0, _ := ret[0].(*product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVendorAndSKU indicates an expected call of FindByVendorAndSKU
func (mr *ProductRepositoryMockRecorder) FindByVendorAndSKU(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVendorAndSKU", reflect.TypeOf((*ProductRepository)(nil).FindByVendorAndSKU), arg0, arg1, arg2)
}

// Upsert mocks base method
func (m *ProductRepository) Upsert(arg0 context.Context, arg1 *product.Product) (product.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(product.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *ProductRepositoryMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*ProductRepository)(nil).Upsert), arg0, arg1)
}
