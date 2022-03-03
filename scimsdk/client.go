package scimsdk

import (
	"scimsdk/internal/service"
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
	trimmedToken := strings.TrimSpace(adminToken)
	client := &Client{opts, trimmedToken}
	return client
}

func (client *Client) GetProvidedURL() string {
	if client.Options != nil {
		return client.Options.APIUrl
	}
	return ""
}

func (client *Client) Users() *UserModule {
	return &UserModule{client: client, service: service.NewUserService(client.token)}
}

func (client *Client) Groups() *GroupModule {
	return &GroupModule{client: client, service: service.NewGroupService(client.token)}
}
