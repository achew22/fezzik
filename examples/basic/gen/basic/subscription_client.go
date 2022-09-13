// Code generated by github.com/inigolabs/fezzik, DO NOT EDIT.

package basic

import (
	"github.com/inigolabs/fezzik/subscription"
)

type SubscriptionClient interface {
	// Updated from examples/basic/operations/operations.graphql:58
	Updated(fn func(out *UpdatedResponse, err error) error) (string, error)

	// Changed from examples/basic/operations/operations.graphql:64
	Changed(
		input *string,
		fn func(out *ChangedResponse, err error) error) (string, error)

	Close() (err error)
}

func NewSubscriptionClient(url string, params map[string]interface{}) SubscriptionClient {
	gql := subscription.NewSubscriptionClient(url)
	gql.WithConnectionParams(params)

	go gql.Run()

	return &gqlSubscriptionClient{gql: gql}
}

type gqlSubscriptionClient struct {
	gql *subscription.SubscriptionClient
}

func (c *gqlSubscriptionClient) Close() (err error) {
	return c.gql.Close()
}