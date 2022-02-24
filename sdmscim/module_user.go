package sdmscim

import (
	"context"
)

type UserModule struct {
	client  *Client
	service *UserService
}

func (module *UserModule) Create(ctx context.Context, user CreateUser) (*User, error) {
	opts := newServiceCreateOptions(&user, module.client.Options.APIUrl)
	return module.service.create(ctx, opts)
}

func (module *UserModule) List(ctx context.Context, paginationOpts *PaginationOptions) *UsersIterator {
	opts := newServiceListOptions(paginationOpts, module.client.Options.APIUrl)
	return module.service.listIterator(ctx, opts)
}

func (module *UserModule) Find(ctx context.Context, id string) (*User, error) {
	opts := newServiceFindOptions(id, module.client.Options.APIUrl)
	return module.service.find(ctx, opts)
}

func (module *UserModule) Replace(ctx context.Context, id string, user ReplaceUser) (*User, error) {
	opts := newServiceReplaceOptions(id, &user, module.client.Options.APIUrl)
	return module.service.replace(ctx, opts)
}

func (module *UserModule) Delete(ctx context.Context, id string) (bool, error) {
	opts := newServiceDeleteOptions(id, module.client.Options.APIUrl)
	return module.service.delete(ctx, opts)
}
