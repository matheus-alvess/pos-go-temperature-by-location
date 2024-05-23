package mock

import (
	"net/http"
)

type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func NewMockHTTPClient(response *http.Response, err error) *MockHTTPClient {
	return &MockHTTPClient{Response: response, Err: err}
}

func (c *MockHTTPClient) Get(url string) (*http.Response, error) {
	return c.Response, c.Err
}
