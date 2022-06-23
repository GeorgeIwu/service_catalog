// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"
	service "service-catalog/service"

	mock "github.com/stretchr/testify/mock"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, attributes
func (_m *ServiceInterface) Create(c context.Context, attributes service.ServiceRequest) (*service.Service, error) {
	ret := _m.Called(c, attributes)

	var r0 *service.Service
	if rf, ok := ret.Get(0).(func(context.Context, service.ServiceRequest) *service.Service); ok {
		r0 = rf(c, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, service.ServiceRequest) error); ok {
		r1 = rf(c, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVersion provides a mock function with given fields: c, serviceID, r
func (_m *ServiceInterface) CreateVersion(c context.Context, serviceID int, r service.VersionRequest) (*service.Service, error) {
	ret := _m.Called(c, serviceID, r)

	var r0 *service.Service
	if rf, ok := ret.Get(0).(func(context.Context, int, service.VersionRequest) *service.Service); ok {
		r0 = rf(c, serviceID, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, service.VersionRequest) error); ok {
		r1 = rf(c, serviceID, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, searchBy, sortBy, pageNumber, itemsPerPage
func (_m *ServiceInterface) Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) ([]*service.Service, error) {
	ret := _m.Called(ctx, searchBy, sortBy, pageNumber, itemsPerPage)

	var r0 []*service.Service
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) []*service.Service); ok {
		r0 = rf(ctx, searchBy, sortBy, pageNumber, itemsPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*service.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int) error); ok {
		r1 = rf(ctx, searchBy, sortBy, pageNumber, itemsPerPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: c, id
func (_m *ServiceInterface) FetchByID(c context.Context, id int) (*service.Service, error) {
	ret := _m.Called(c, id)

	var r0 *service.Service
	if rf, ok := ret.Get(0).(func(context.Context, int) *service.Service); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchVersions provides a mock function with given fields: c, serviceID
func (_m *ServiceInterface) FetchVersions(c context.Context, serviceID int) ([]*service.Version, error) {
	ret := _m.Called(c, serviceID)

	var r0 []*service.Version
	if rf, ok := ret.Get(0).(func(context.Context, int) []*service.Version); ok {
		r0 = rf(c, serviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*service.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(c, serviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: c, serviceID
func (_m *ServiceInterface) Remove(c context.Context, serviceID int) error {
	ret := _m.Called(c, serviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(c, serviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: c, serviceID, r
func (_m *ServiceInterface) Update(c context.Context, serviceID int, r service.ServiceRequest) (*service.Service, error) {
	ret := _m.Called(c, serviceID, r)

	var r0 *service.Service
	if rf, ok := ret.Get(0).(func(context.Context, int, service.ServiceRequest) *service.Service); ok {
		r0 = rf(c, serviceID, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, service.ServiceRequest) error); ok {
		r1 = rf(c, serviceID, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceInterface creates a new instance of ServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceInterface(t mockConstructorTestingTNewServiceInterface) *ServiceInterface {
	mock := &ServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}