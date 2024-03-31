package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// HTTPDoer represents any struct that is capable of executing `Do`
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type GQLClient struct {
	endpoint   string
	httpClient HTTPDoer
	debug      bool
}

type Option func(c *GQLClient)

func WithDebug() Option {
	return func(c *GQLClient) {
		c.debug = true
	}
}

func NewGQLClient(endpoint string, httpClient HTTPDoer, opts ...Option) *GQLClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	client := &GQLClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}

	for _, o := range opts {
		o(client)
	}

	return client
}

func (c *GQLClient) SetHttpClient(httpClient HTTPDoer) {
	c.httpClient = httpClient
}

func (c *GQLClient) GetHttpClient() HTTPDoer {
	return c.httpClient
}

func (c *GQLClient) GetEndpoint() string {
	return c.endpoint
}

func (c *GQLClient) Query(ctx context.Context, req *GQLRequest, res *GQLResponse) error {
	var body bytes.Buffer
	encoder := json.NewEncoder(&body)
	err := encoder.Encode(req)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, c.endpoint, &body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json; charset=utf-8")
	request = request.WithContext(ctx)
	request.Close = true
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(res)
	if err != nil {
		if c.debug {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return fmt.Errorf("Code: %d. Cannot read body.", response.StatusCode)
			}

			return fmt.Errorf("Code: %d. Body: %s", response.StatusCode, body)
		}

		return fmt.Errorf("Code: %d. Invalid graphql response.", response.StatusCode)
	}

	return nil
}
