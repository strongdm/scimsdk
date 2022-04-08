package scimsdk

import (
	"strings"

	"github.com/strongdm/scimsdk/internal/api"
	"github.com/strongdm/scimsdk/internal/module"
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
	token   string
	options *ClientOptions
}

func NewClient(adminToken string, opts *ClientOptions) Client {
	trimmedToken := strings.TrimSpace(adminToken)
	client := &clientImpl{trimmedToken, opts}
	return client
}

func (client *clientImpl) Users() UserModule {
	return module.NewUserModule(service.NewUserService(api.NewAPI(), client.getToken()), client.GetProvidedURL())
}

func (client *clientImpl) Groups() GroupModule {
	return module.NewGroupModule(service.NewGroupService(api.NewAPI(), client.getToken()), client.GetProvidedURL())
}

func (client *clientImpl) GetProvidedURL() string {
	if client.options != nil {
		return client.options.APIUrl
	}
	return ""
}

func (client *clientImpl) getToken() string {
	return client.token
}
