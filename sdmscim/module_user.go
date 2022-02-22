package sdmscim

import (
	"context"
)

type UserModule struct {
	client *Client
}

func (module UserModule) List(ctx context.Context, opts *ModuleListOptions) *UsersIterator {
	if opts == nil {
		opts = &ModuleListOptions{}
	}
	if module.client.Options.APIUrl != "" {
		opts.SetBaseAPIURL(module.client.Options.APIUrl)
	}
	service := newUserService(module.client.adminToken, ctx)
	iterator := newUsersIterator(service.list, opts)
	return iterator
}
