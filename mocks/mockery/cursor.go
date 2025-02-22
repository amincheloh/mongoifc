// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	bson "go.mongodb.org/mongo-driver/bson"

	mock "github.com/stretchr/testify/mock"
)

// Cursor is an autogenerated mock type for the Cursor type
type Cursor struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx, results
func (_m *Cursor) All(ctx context.Context, results interface{}) error {
	ret := _m.Called(ctx, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields: ctx
func (_m *Cursor) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Current provides a mock function with given fields:
func (_m *Cursor) Current() bson.Raw {
	ret := _m.Called()

	var r0 bson.Raw
	if rf, ok := ret.Get(0).(func() bson.Raw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	return r0
}

// Decode provides a mock function with given fields: val
func (_m *Cursor) Decode(val interface{}) error {
	ret := _m.Called(val)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Err provides a mock function with given fields:
func (_m *Cursor) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Cursor) ID() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Next provides a mock function with given fields: ctx
func (_m *Cursor) Next(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemainingBatchLength provides a mock function with given fields:
func (_m *Cursor) RemainingBatchLength() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SetBatchSize provides a mock function with given fields: batchSize
func (_m *Cursor) SetBatchSize(batchSize int32) {
	_m.Called(batchSize)
}

// TryNext provides a mock function with given fields: ctx
func (_m *Cursor) TryNext(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewCursor creates a new instance of Cursor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCursor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cursor {
	mock := &Cursor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
