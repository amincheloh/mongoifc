// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	bson "go.mongodb.org/mongo-driver/bson"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"

	mongoifc "github.com/sv-tools/mongoifc"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// IndexView is an autogenerated mock type for the IndexView type
type IndexView struct {
	mock.Mock
}

// CreateMany provides a mock function with given fields: ctx, models, opts
func (_m *IndexView) CreateMany(ctx context.Context, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, models)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []mongo.IndexModel, ...*options.CreateIndexesOptions) ([]string, error)); ok {
		return rf(ctx, models, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []mongo.IndexModel, ...*options.CreateIndexesOptions) []string); ok {
		r0 = rf(ctx, models, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []mongo.IndexModel, ...*options.CreateIndexesOptions) error); ok {
		r1 = rf(ctx, models, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOne provides a mock function with given fields: ctx, model, opts
func (_m *IndexView) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, model)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, mongo.IndexModel, ...*options.CreateIndexesOptions) (string, error)); ok {
		return rf(ctx, model, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, mongo.IndexModel, ...*options.CreateIndexesOptions) string); ok {
		r0 = rf(ctx, model, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, mongo.IndexModel, ...*options.CreateIndexesOptions) error); ok {
		r1 = rf(ctx, model, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DropAll provides a mock function with given fields: ctx, opts
func (_m *IndexView) DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bson.Raw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*options.DropIndexesOptions) (bson.Raw, error)); ok {
		return rf(ctx, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...*options.DropIndexesOptions) bson.Raw); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...*options.DropIndexesOptions) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DropOne provides a mock function with given fields: ctx, name, opts
func (_m *IndexView) DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bson.Raw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...*options.DropIndexesOptions) (bson.Raw, error)); ok {
		return rf(ctx, name, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...*options.DropIndexesOptions) bson.Raw); ok {
		r0 = rf(ctx, name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...*options.DropIndexesOptions) error); ok {
		r1 = rf(ctx, name, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *IndexView) List(ctx context.Context, opts ...*options.ListIndexesOptions) (mongoifc.Cursor, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongoifc.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*options.ListIndexesOptions) (mongoifc.Cursor, error)); ok {
		return rf(ctx, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...*options.ListIndexesOptions) mongoifc.Cursor); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...*options.ListIndexesOptions) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSpecifications provides a mock function with given fields: ctx, opts
func (_m *IndexView) ListSpecifications(ctx context.Context, opts ...*options.ListIndexesOptions) ([]*mongo.IndexSpecification, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*mongo.IndexSpecification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*options.ListIndexesOptions) ([]*mongo.IndexSpecification, error)); ok {
		return rf(ctx, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...*options.ListIndexesOptions) []*mongo.IndexSpecification); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*mongo.IndexSpecification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...*options.ListIndexesOptions) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIndexView creates a new instance of IndexView. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndexView(t interface {
	mock.TestingT
	Cleanup(func())
}) *IndexView {
	mock := &IndexView{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
