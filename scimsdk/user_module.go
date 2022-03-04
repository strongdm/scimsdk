package scimsdk

import (
	"context"

	"github.com/strongdm/scimsdk/internal/service"
)

type UserModule struct {
	client  *Client
	service *service.UserService
}

func (module *UserModule) Create(ctx context.Context, user CreateUser) (*User, error) {
	body := convertPorcelainToCreateUserRequest(&user)
	opts := newServiceCreateOptions(body, module.client.GetProvidedURL())
	response, err := module.service.Create(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *UserModule) List(ctx context.Context, paginationOpts *PaginationOptions) *UsersIterator {
	return newUsersIterator(module.iteratorMiddleware(ctx), paginationOpts)
}

func (module *UserModule) Find(ctx context.Context, id string) (*User, error) {
	opts := newServiceFindOptions(id, module.client.GetProvidedURL())
	response, err := module.service.Find(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *UserModule) Replace(ctx context.Context, id string, user ReplaceUser) (*User, error) {
	body := convertPorcelainToReplaceUserRequest(id, &user)
	opts := newServiceReplaceOptions(id, body, module.client.GetProvidedURL())
	response, err := module.service.Replace(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *UserModule) Update(ctx context.Context, id string, updateUser UpdateUser) (bool, error) {
	body := convertPorcelainToUpdateUserRequest(updateUser)
	opts := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	return module.service.Update(ctx, opts)
}

func (module *UserModule) Delete(ctx context.Context, id string) (bool, error) {
	opts := newServiceDeleteOptions(id, module.client.GetProvidedURL())
	return module.service.Delete(ctx, opts)
}

func (module *UserModule) iteratorMiddleware(ctx context.Context) func(opts *PaginationOptions) ([]*User, bool, error) {
	return func(opts *PaginationOptions) ([]*User, bool, error) {
		listOpts := newServiceListOptions(opts, module.client.GetProvidedURL())
		response, haveNextPage, err := module.service.List(ctx, listOpts)
		if err != nil {
			return nil, false, err
		}
		users := convertUserResponseListToPorcelain(response)
		return users, haveNextPage, nil
	}
}
