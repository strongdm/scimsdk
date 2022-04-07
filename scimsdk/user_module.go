package scimsdk

import (
	"context"

	"github.com/strongdm/scimsdk/internal/service"
)

type UserModule interface {
	Create(ctx context.Context, user CreateUser) (*User, error)
	List(ctx context.Context, paginationOpts *PaginationOptions) UserIterator
	Find(ctx context.Context, id string) (*User, error)
	Replace(ctx context.Context, id string, user ReplaceUser) (*User, error)
	Update(ctx context.Context, id string, updateUser UpdateUser) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type userModuleImpl struct {
	client  Client
	service service.UserService
}

func NewUserModule(client Client, service service.UserService) UserModule {
	return &userModuleImpl{client, service}
}

func (module *userModuleImpl) Create(ctx context.Context, user CreateUser) (*User, error) {
	body, err := convertPorcelainToCreateUserRequest(&user)
	if err != nil {
		return nil, err
	}
	opts := newServiceCreateOptions(body, module.client.GetProvidedURL())
	response, err := module.service.Create(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *userModuleImpl) List(ctx context.Context, paginationOpts *PaginationOptions) UserIterator {
	return newUsersIterator(module.iteratorMiddleware(ctx), paginationOpts)
}

func (module *userModuleImpl) Find(ctx context.Context, id string) (*User, error) {
	opts, err := newServiceFindOptions(id, module.client.GetProvidedURL())
	if err != nil {
		return nil, err
	}
	response, err := module.service.Find(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *userModuleImpl) Replace(ctx context.Context, id string, user ReplaceUser) (*User, error) {
	body, err := convertPorcelainToReplaceUserRequest(id, &user)
	if err != nil {
		return nil, err
	}
	opts, err := newServiceReplaceOptions(id, body, module.client.GetProvidedURL())
	if err != nil {
		return nil, err
	}
	response, err := module.service.Replace(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *userModuleImpl) Update(ctx context.Context, id string, updateUser UpdateUser) (bool, error) {
	body := convertPorcelainToUpdateUserRequest(updateUser)
	opts, err := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *userModuleImpl) Delete(ctx context.Context, id string) (bool, error) {
	opts, err := newServiceDeleteOptions(id, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Delete(ctx, opts)
}

func (module *userModuleImpl) iteratorMiddleware(ctx context.Context) listUsersOperationFunc {
	return func(opts *PaginationOptions) ([]*User, bool, error) {
		listOpts, err := newServiceListOptions(opts, module.client.GetProvidedURL())
		if err != nil {
			return nil, false, err
		}
		response, haveNextPage, err := module.service.List(ctx, listOpts)
		if err != nil {
			return nil, false, err
		}
		users := convertUserResponseListToPorcelain(response)
		return users, haveNextPage, nil
	}
}
