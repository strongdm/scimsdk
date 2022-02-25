package sdmscim

import (
	"context"
)

type UserModule struct {
	client  *Client
	service *UserService
}

func (module *UserModule) Create(ctx context.Context, user CreateUser) (*User, error) {
	body := convertPorcelainToCreateUserRequest(&user)
	opts := newServiceCreateOptions(body, module.client.Options.APIUrl)
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
	body := convertPorcelainToReplaceUserRequest(id, &user)
	opts := newServiceReplaceOptions(id, body, module.client.Options.APIUrl)
	return module.service.replace(ctx, opts)
}

func (module *UserModule) Update(ctx context.Context, id string, updateUser UpdateUser) (bool, error) {
	body := convertPorcelainToUpdateUserRequest(updateUser)
	opts := newServiceUpdateOptions(id, body, module.client.Options.APIUrl)
	return module.service.update(ctx, opts)
}

func (module *UserModule) Delete(ctx context.Context, id string) (bool, error) {
	opts := newServiceDeleteOptions(id, module.client.Options.APIUrl)
	return module.service.delete(ctx, opts)
}
