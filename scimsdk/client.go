package scimsdk

import (
	"strings"

	"github.com/strongdm/scimsdk/internal/api"
	"github.com/strongdm/scimsdk/internal/service"
)

type Client interface {
	Users() UserModule
	Groups() GroupModule
	GetProvidedURL() string
}

type ClientOptions struct {
	APIUrl string
}

type clientImpl struct {
	options *ClientOptions
	token   string
}

func NewClient(adminToken string, opts *ClientOptions) Client {
	trimmedToken := strings.TrimSpace(adminToken)
	client := &clientImpl{opts, trimmedToken}
	return client
}

func (client *clientImpl) GetProvidedURL() string {
	if client.options != nil {
		return client.options.APIUrl
	}
	return ""
}

func (client *clientImpl) Users() UserModule {
	return NewUserModule(client, service.NewUserService(api.NewAPI(), client.token))
}

func (client *clientImpl) Groups() GroupModule {
	return NewGroupModule(client, service.NewGroupService(api.NewAPI(), client.token))
}
