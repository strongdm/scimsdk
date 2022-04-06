package scimsdk

import (
	"strings"

	"github.com/strongdm/scimsdk/internal/api"
	"github.com/strongdm/scimsdk/internal/service"
)

type IClient interface {
	Users() IUserModule
	Groups() IGroupModule
	GetProvidedURL() string
}

type ClientOptions struct {
	APIUrl string
}

type Client struct {
	options *ClientOptions
	token   string
}

func NewClient(adminToken string, opts *ClientOptions) IClient {
	trimmedToken := strings.TrimSpace(adminToken)
	client := &Client{opts, trimmedToken}
	return client
}

func (client *Client) GetProvidedURL() string {
	if client.options != nil {
		return client.options.APIUrl
	}
	return ""
}

func (client *Client) Users() IUserModule {
	return NewUserModule(client, service.NewUserService(api.NewAPI(), client.token))
}

func (client *Client) Groups() IGroupModule {
	return NewGroupModule(client, service.NewGroupService(api.NewAPI(), client.token))
}
