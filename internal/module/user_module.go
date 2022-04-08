package module

import (
	"context"

	"github.com/strongdm/scimsdk/internal/service"
	"github.com/strongdm/scimsdk/models"
)

type userModuleImpl struct {
	service     service.UserService
	providedURL string
}

func NewUserModule(service service.UserService, providedURL string) *userModuleImpl {
	return &userModuleImpl{service, providedURL}
}

func (module *userModuleImpl) Create(ctx context.Context, user models.CreateUser) (*models.User, error) {
	body, err := convertPorcelainToCreateUserRequest(&user)
	if err != nil {
		return nil, err
	}
	opts := newServiceCreateOptions(body, module.providedURL)
	response, err := module.service.Create(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *userModuleImpl) List(ctx context.Context, paginationOpts *models.PaginationOptions) models.Iterator[models.User] {
	return newIterator(module.iteratorMiddleware(ctx), paginationOpts)
}

func (module *userModuleImpl) Find(ctx context.Context, id string) (*models.User, error) {
	opts, err := newServiceFindOptions(id, module.providedURL)
	if err != nil {
		return nil, err
	}
	response, err := module.service.Find(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *userModuleImpl) Replace(ctx context.Context, id string, user models.ReplaceUser) (*models.User, error) {
	body, err := convertPorcelainToReplaceUserRequest(id, &user)
	if err != nil {
		return nil, err
	}
	opts, err := newServiceReplaceOptions(id, body, module.providedURL)
	if err != nil {
		return nil, err
	}
	response, err := module.service.Replace(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(response), nil
}

func (module *userModuleImpl) Update(ctx context.Context, id string, updateUser models.UpdateUser) (bool, error) {
	body := convertPorcelainToUpdateUserRequest(updateUser)
	opts, err := newServiceUpdateOptions(id, body, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *userModuleImpl) Delete(ctx context.Context, id string) (bool, error) {
	opts, err := newServiceDeleteOptions(id, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Delete(ctx, opts)
}

func (module *userModuleImpl) iteratorMiddleware(ctx context.Context) iteratorFetchFunc[models.User] {
	return func(opts *models.PaginationOptions) ([]*models.User, bool, error) {
		listOpts, err := newServiceListOptions(opts, module.providedURL)
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
