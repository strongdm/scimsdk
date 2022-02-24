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

func NewClient(adminToken string, opts *ClientOptions) *Client {
	if opts == nil {
		opts = getDefaultClientOptions()
	}
	trimmedToken := strings.TrimSpace(adminToken)
	client := &Client{opts, trimmedToken}
	return client
}

func getDefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		APIUrl: defaultAPIURL,
	}
}

func (client *Client) Users() *UserModule {
	return &UserModule{client: client, service: newUserService(client.token)}
}
