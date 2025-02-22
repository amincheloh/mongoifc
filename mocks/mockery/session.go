// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	bson "go.mongodb.org/mongo-driver/bson"

	mock "github.com/stretchr/testify/mock"

	mongoifc "github.com/sv-tools/mongoifc"

	options "go.mongodb.org/mongo-driver/mongo/options"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Session is an autogenerated mock type for the Session type
type Session struct {
	mock.Mock
}

// AbortTransaction provides a mock function with given fields: ctx
func (_m *Session) AbortTransaction(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdvanceClusterTime provides a mock function with given fields: _a0
func (_m *Session) AdvanceClusterTime(_a0 bson.Raw) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(bson.Raw) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdvanceOperationTime provides a mock function with given fields: _a0
func (_m *Session) AdvanceOperationTime(_a0 *primitive.Timestamp) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*primitive.Timestamp) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client provides a mock function with given fields:
func (_m *Session) Client() mongoifc.Client {
	ret := _m.Called()

	var r0 mongoifc.Client
	if rf, ok := ret.Get(0).(func() mongoifc.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Client)
		}
	}

	return r0
}

// ClusterTime provides a mock function with given fields:
func (_m *Session) ClusterTime() bson.Raw {
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

// CommitTransaction provides a mock function with given fields: ctx
func (_m *Session) CommitTransaction(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EndSession provides a mock function with given fields: ctx
func (_m *Session) EndSession(ctx context.Context) {
	_m.Called(ctx)
}

// ID provides a mock function with given fields:
func (_m *Session) ID() bson.Raw {
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

// OperationTime provides a mock function with given fields:
func (_m *Session) OperationTime() *primitive.Timestamp {
	ret := _m.Called()

	var r0 *primitive.Timestamp
	if rf, ok := ret.Get(0).(func() *primitive.Timestamp); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*primitive.Timestamp)
		}
	}

	return r0
}

// StartTransaction provides a mock function with given fields: opts
func (_m *Session) StartTransaction(opts ...*options.TransactionOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*options.TransactionOptions) error); ok {
		r0 = rf(opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithTransaction provides a mock function with given fields: ctx, fn, opts
func (_m *Session) WithTransaction(ctx context.Context, fn func(mongoifc.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, fn)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, func(mongoifc.SessionContext) (interface{}, error), ...*options.TransactionOptions) (interface{}, error)); ok {
		return rf(ctx, fn, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, func(mongoifc.SessionContext) (interface{}, error), ...*options.TransactionOptions) interface{}); ok {
		r0 = rf(ctx, fn, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, func(mongoifc.SessionContext) (interface{}, error), ...*options.TransactionOptions) error); ok {
		r1 = rf(ctx, fn, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSession creates a new instance of Session. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSession(t interface {
	mock.TestingT
	Cleanup(func())
}) *Session {
	mock := &Session{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
