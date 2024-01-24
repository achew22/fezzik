// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	basic "github.com/inigolabs/fezzik/examples/basic/gen/basic"

	mock "github.com/stretchr/testify/mock"

	types "github.com/inigolabs/fezzik/examples/basic/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// InterfaceGetCreated provides a mock function with given fields: ctx
func (_m *Client) InterfaceGetCreated(ctx context.Context) (*basic.InterfaceGetCreatedResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.InterfaceGetCreatedResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.InterfaceGetCreatedResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.InterfaceGetCreatedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InterfaceGetCreatedNoTypename provides a mock function with given fields: ctx
func (_m *Client) InterfaceGetCreatedNoTypename(ctx context.Context) (*basic.InterfaceGetCreatedNoTypenameResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.InterfaceGetCreatedNoTypenameResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.InterfaceGetCreatedNoTypenameResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.InterfaceGetCreatedNoTypenameResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OneAdd provides a mock function with given fields: ctx, input
func (_m *Client) OneAdd(ctx context.Context, input *basic.OneInput) (*basic.OneAddResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *basic.OneAddResponse
	if rf, ok := ret.Get(0).(func(context.Context, *basic.OneInput) *basic.OneAddResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.OneAddResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *basic.OneInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OneAllTypes provides a mock function with given fields: ctx
func (_m *Client) OneAllTypes(ctx context.Context) (*basic.OneAllTypesResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.OneAllTypesResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.OneAllTypesResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.OneAllTypesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OneWithAlias provides a mock function with given fields: ctx
func (_m *Client) OneWithAlias(ctx context.Context) (*basic.OneWithAliasResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.OneWithAliasResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.OneWithAliasResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.OneWithAliasResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OneWithSubSelections provides a mock function with given fields: ctx
func (_m *Client) OneWithSubSelections(ctx context.Context) (*basic.OneWithSubSelectionsResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.OneWithSubSelectionsResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.OneWithSubSelectionsResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.OneWithSubSelectionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryWithInputs provides a mock function with given fields: ctx, inputOne, inputTwo
func (_m *Client) QueryWithInputs(ctx context.Context, inputOne *string, inputTwo *string) (*basic.QueryWithInputsResponse, error) {
	ret := _m.Called(ctx, inputOne, inputTwo)

	var r0 *basic.QueryWithInputsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *string, *string) *basic.QueryWithInputsResponse); ok {
		r0 = rf(ctx, inputOne, inputTwo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.QueryWithInputsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *string, *string) error); ok {
		r1 = rf(ctx, inputOne, inputTwo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TwoAdd provides a mock function with given fields: ctx, input
func (_m *Client) TwoAdd(ctx context.Context, input *types.TwoInput) (*basic.TwoAddResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *basic.TwoAddResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.TwoInput) *basic.TwoAddResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.TwoAddResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.TwoInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnionSearchResult provides a mock function with given fields: ctx
func (_m *Client) UnionSearchResult(ctx context.Context) (*basic.UnionSearchResultResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.UnionSearchResultResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.UnionSearchResultResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.UnionSearchResultResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnionSearchResultInLine provides a mock function with given fields: ctx
func (_m *Client) UnionSearchResultInLine(ctx context.Context) (*basic.UnionSearchResultInLineResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.UnionSearchResultInLineResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.UnionSearchResultInLineResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.UnionSearchResultInLineResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnionSearchResultNoTypename provides a mock function with given fields: ctx
func (_m *Client) UnionSearchResultNoTypename(ctx context.Context) (*basic.UnionSearchResultNoTypenameResponse, error) {
	ret := _m.Called(ctx)

	var r0 *basic.UnionSearchResultNoTypenameResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.UnionSearchResultNoTypenameResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.UnionSearchResultNoTypenameResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
