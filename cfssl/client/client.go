package client

import (
	"github.com/hashicorp/go-cleanhttp"
	"net/http"
	"time"
)

const (
	DefaultTimeout = 15
)

type Client struct {
	httpClient *http.Client
	uri        string
	headers    map[string]string
}

// Send a http HEAD request to the cfssl uri to check so that the server is available.
func (client Client) Ping() error {
	res, err := client.httpClient.Head(client.uri)
	if err == nil && res != nil {
		return nil
	}
	return err
}

// Creates a new client object and sets the uri, default headers and creates a new httpClient to use
// for requests to the cfssl API.
func New(uri string, headers map[string]string) *Client {
	client := Client{}
	client.uri = uri
	client.headers = headers
	client.httpClient = cleanhttp.DefaultClient()
	client.httpClient.Timeout = time.Second * DefaultTimeout
	return &client
}
