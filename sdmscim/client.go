package sdmscim

import (
	"strings"
)

type ClientOptions struct {
	APIUrl string
}

type Client struct {
	Users      *UserModule
	Options    *ClientOptions
	adminToken string
}

var defaultClientOptions *ClientOptions = &ClientOptions{
	APIUrl: defaultAPIURL,
}

func NewClient(adminToken string, opts *ClientOptions) *Client {
	if opts == nil {
		opts = defaultClientOptions
	}
	trimmedToken := strings.TrimSpace(adminToken)
	client := &Client{adminToken: trimmedToken, Options: opts}
	client.Users = &UserModule{client: client}
	return client
}
