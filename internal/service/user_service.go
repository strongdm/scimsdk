package service

import (
	"context"

	"github.com/strongdm/scimsdk/internal/api"
)

type UserService struct {
	token string
}

const usersAPIPathname = "Users"

func NewUserService(token string) *UserService {
	return &UserService{token: token}
}

func (service *UserService) Create(ctx context.Context, opts *CreateOptions) (*UserResponse, error) {
	response, err := api.Create(ctx, usersAPIPathname, service.token, newAPICreateOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalUserResponse(response.Body)
}

func (service *UserService) List(ctx context.Context, opts *ListOptions) ([]UserResponse, bool, error) {
	response, err := api.List(ctx, usersAPIPathname, service.token, newAPIListOptions(opts))
	if err != nil {
		return nil, false, err
	}
	unmarshedResponse, err := unmarshalUserPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	pageSize := api.GetDefaultAPIPageSize()
	if opts.PageSize != 0 {
		pageSize = opts.PageSize
	}
	return unmarshedResponse.Resources, len(unmarshedResponse.Resources) >= pageSize, nil
}

func (service *UserService) Find(ctx context.Context, opts *FindOptions) (*UserResponse, error) {
	response, err := api.Find(ctx, usersAPIPathname, service.token, newAPIFindOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalUserResponse(response.Body)
}

func (service *UserService) Replace(ctx context.Context, opts *ReplaceOptions) (*UserResponse, error) {
	response, err := api.Replace(ctx, usersAPIPathname, service.token, newAPIReplaceOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalUserResponse(response.Body)
}

func (service *UserService) Update(ctx context.Context, opts *UpdateOptions) (bool, error) {
	_, err := api.Update(ctx, usersAPIPathname, service.token, newAPIUpdateOptions(opts))
	return err == nil, err
}

func (service *UserService) Delete(ctx context.Context, opts *DeleteOptions) (bool, error) {
	_, err := api.Delete(ctx, usersAPIPathname, service.token, newAPIDeleteOptions(opts))
	if err != nil {
		return false, err
	}
	return true, nil
}
