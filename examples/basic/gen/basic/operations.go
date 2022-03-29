package basic

import (
	"context"

	"github.com/inigolabs/fezzik/client"
)

type OneAllTypesOne struct {
	OneInt             *int
	OneIntMust         int
	OneIntList         *[]*int
	OneIntMustList     *[]int
	OneIntMustListMust []int
}

// OneAllTypesResponse response type for OneAllTypes
type OneAllTypesResponse struct {
	One *OneAllTypesOne
}

// OneAllTypes from examples/basic/operations/operations.graphql:2
func (c *gqlclient) OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error) {

	var oneAllTypesOperation string = `
	query OneAllTypes {
		one {
			oneInt
			oneIntMust
			oneIntList
			oneIntMustList
			oneIntMustListMust
		}
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "OneAllTypes",
		Query:         oneAllTypesOperation,
		Variables:     map[string]interface{}{},
	}

	var gqldata OneAllTypesResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

type OneWithSubSelectionsOne struct {
	Two *OneWithSubSelectionsOneTwo
}

type OneWithSubSelectionsOneTwo struct {
	TwoInt *int
	TwoStr *string
	Three  *OneWithSubSelectionsOneTwoThree
}

type OneWithSubSelectionsOneTwoThree struct {
	ThreeInt *int
}

// OneWithSubSelectionsResponse response type for OneWithSubSelections
type OneWithSubSelectionsResponse struct {
	One *OneWithSubSelectionsOne
}

// OneWithSubSelections from examples/basic/operations/operations.graphql:13
func (c *gqlclient) OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error) {

	var oneWithSubSelectionsOperation string = `
	query OneWithSubSelections {
		one {
			two {
				twoInt
				twoStr
				three {
					threeInt
				}
			}
		}
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "OneWithSubSelections",
		Query:         oneWithSubSelectionsOperation,
		Variables:     map[string]interface{}{},
	}

	var gqldata OneWithSubSelectionsResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

type QueryWithInputsOne struct {
	OneInt *int
}

type QueryWithInputsTwo struct {
	TwoInt *int
}

// QueryWithInputsResponse response type for QueryWithInputs
type QueryWithInputsResponse struct {
	One *QueryWithInputsOne
	Two *QueryWithInputsTwo
}

// QueryWithInputs from examples/basic/operations/operations.graphql:27
func (c *gqlclient) QueryWithInputs(ctx context.Context,
	inputOne *string,
	inputTwo *string,
) (*QueryWithInputsResponse, error) {

	var queryWithInputsOperation string = `
	query QueryWithInputs($input_one : String, $input_two : String) {
		one(input: $input_one) {
			oneInt
		}
		two(input: $input_two) {
			twoInt
		}
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "QueryWithInputs",
		Query:         queryWithInputsOperation,
		Variables: map[string]interface{}{
			"input_one": inputOne,
			"input_two": inputTwo,
		},
	}

	var gqldata QueryWithInputsResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

// OneAddResponse response type for OneAdd
type OneAddResponse struct {
	OneAdd *string
}

// OneAdd from examples/basic/operations/operations.graphql:38
func (c *gqlclient) OneAdd(ctx context.Context,
	input *OneInput,
) (*OneAddResponse, error) {

	var oneAddOperation string = `
	mutation OneAdd($input : OneInput) {
		oneAdd(input: $input)
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "OneAdd",
		Query:         oneAddOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	var gqldata OneAddResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}
