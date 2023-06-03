// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "gorm-test/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// BookWriter is an autogenerated mock type for the BookWriter type
type BookWriter struct {
	mock.Mock
}

// Insert provides a mock function with given fields: book
func (_m *BookWriter) Insert(book *models.Book) error {
	ret := _m.Called(book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Book) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBookWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookWriter creates a new instance of BookWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookWriter(t mockConstructorTestingTNewBookWriter) *BookWriter {
	mock := &BookWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
