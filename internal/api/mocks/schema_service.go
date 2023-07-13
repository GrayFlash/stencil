// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	schema "github.com/raystack/stencil/core/schema"
	mock "github.com/stretchr/testify/mock"
)

// SchemaService is an autogenerated mock type for the SchemaService type
type SchemaService struct {
	mock.Mock
}

// CheckCompatibility provides a mock function with given fields: ctx, nsName, schemaName, compatibility, data
func (_m *SchemaService) CheckCompatibility(ctx context.Context, nsName string, schemaName string, compatibility string, data []byte) error {
	ret := _m.Called(ctx, nsName, schemaName, compatibility, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []byte) error); ok {
		r0 = rf(ctx, nsName, schemaName, compatibility, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, nsName, schemaName, metadata, data
func (_m *SchemaService) Create(ctx context.Context, nsName string, schemaName string, metadata *schema.Metadata, data []byte) (schema.SchemaInfo, error) {
	ret := _m.Called(ctx, nsName, schemaName, metadata, data)

	var r0 schema.SchemaInfo
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *schema.Metadata, []byte) schema.SchemaInfo); ok {
		r0 = rf(ctx, nsName, schemaName, metadata, data)
	} else {
		r0 = ret.Get(0).(schema.SchemaInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *schema.Metadata, []byte) error); ok {
		r1 = rf(ctx, nsName, schemaName, metadata, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, namespace, schemaName
func (_m *SchemaService) Delete(ctx context.Context, namespace string, schemaName string) error {
	ret := _m.Called(ctx, namespace, schemaName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, namespace, schemaName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVersion provides a mock function with given fields: ctx, namespace, schemaName, version
func (_m *SchemaService) DeleteVersion(ctx context.Context, namespace string, schemaName string, version int32) error {
	ret := _m.Called(ctx, namespace, schemaName, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int32) error); ok {
		r0 = rf(ctx, namespace, schemaName, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, namespace, schemaName, version
func (_m *SchemaService) Get(ctx context.Context, namespace string, schemaName string, version int32) (*schema.Metadata, []byte, error) {
	ret := _m.Called(ctx, namespace, schemaName, version)

	var r0 *schema.Metadata
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int32) *schema.Metadata); ok {
		r0 = rf(ctx, namespace, schemaName, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.Metadata)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int32) []byte); ok {
		r1 = rf(ctx, namespace, schemaName, version)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, int32) error); ok {
		r2 = rf(ctx, namespace, schemaName, version)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLatest provides a mock function with given fields: ctx, namespace, schemaName
func (_m *SchemaService) GetLatest(ctx context.Context, namespace string, schemaName string) (*schema.Metadata, []byte, error) {
	ret := _m.Called(ctx, namespace, schemaName)

	var r0 *schema.Metadata
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *schema.Metadata); ok {
		r0 = rf(ctx, namespace, schemaName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.Metadata)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string, string) []byte); ok {
		r1 = rf(ctx, namespace, schemaName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, namespace, schemaName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetMetadata provides a mock function with given fields: ctx, namespace, schemaName
func (_m *SchemaService) GetMetadata(ctx context.Context, namespace string, schemaName string) (*schema.Metadata, error) {
	ret := _m.Called(ctx, namespace, schemaName)

	var r0 *schema.Metadata
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *schema.Metadata); ok {
		r0 = rf(ctx, namespace, schemaName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, schemaName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, namespaceID
func (_m *SchemaService) List(ctx context.Context, namespaceID string) ([]schema.Schema, error) {
	ret := _m.Called(ctx, namespaceID)

	var r0 []schema.Schema
	if rf, ok := ret.Get(0).(func(context.Context, string) []schema.Schema); ok {
		r0 = rf(ctx, namespaceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]schema.Schema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, namespaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVersions provides a mock function with given fields: ctx, namespaceID, schemaName
func (_m *SchemaService) ListVersions(ctx context.Context, namespaceID string, schemaName string) ([]int32, error) {
	ret := _m.Called(ctx, namespaceID, schemaName)

	var r0 []int32
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []int32); ok {
		r0 = rf(ctx, namespaceID, schemaName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespaceID, schemaName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMetadata provides a mock function with given fields: ctx, namespace, schemaName, meta
func (_m *SchemaService) UpdateMetadata(ctx context.Context, namespace string, schemaName string, meta *schema.Metadata) (*schema.Metadata, error) {
	ret := _m.Called(ctx, namespace, schemaName, meta)

	var r0 *schema.Metadata
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *schema.Metadata) *schema.Metadata); ok {
		r0 = rf(ctx, namespace, schemaName, meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *schema.Metadata) error); ok {
		r1 = rf(ctx, namespace, schemaName, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSchemaService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSchemaService creates a new instance of SchemaService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSchemaService(t mockConstructorTestingTNewSchemaService) *SchemaService {
	mock := &SchemaService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}