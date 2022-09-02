// Code generated by github.com/inigolabs/fezzik, DO NOT EDIT.

package github

import (
	subscription "github.com/hasura/go-graphql-client"
)

type SubscriptionClient interface {
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
