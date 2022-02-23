package sdmscim

import (
	"strings"
)

type ClientOptions struct {
	APIUrl string
}

type Client struct {
	Options *ClientOptions
	token   string
}

var defaultClientOptions *ClientOptions = &ClientOptions{
	APIUrl: defaultAPIURL,
}

func NewClient(adminToken string, opts *ClientOptions) *Client {
	if opts == nil {
		opts = defaultClientOptions
	}
	trimmedToken := strings.TrimSpace(adminToken)
	client := &Client{opts, trimmedToken}
	return client
}

func (client *Client) Users() *UserModule {
	return &UserModule{client: client, service: newUserService(client.token)}
}
