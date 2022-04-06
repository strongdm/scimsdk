package service

import (
	"context"

	"github.com/strongdm/scimsdk/internal/api"
)

type IUserService interface {
	Create(ctx context.Context, opts *CreateOptions) (*UserResponse, error)
	List(ctx context.Context, opts *ListOptions) ([]*UserResponse, bool, error)
	Find(ctx context.Context, opts *FindOptions) (*UserResponse, error)
	Replace(ctx context.Context, opts *ReplaceOptions) (*UserResponse, error)
	Update(ctx context.Context, opts *UpdateOptions) (bool, error)
	Delete(ctx context.Context, opts *DeleteOptions) (bool, error)
}

type UserService struct {
	client api.IAPI
	token  string
}

const usersAPIPathname = "Users"

func NewUserService(api api.IAPI, token string) IUserService {
	return &UserService{api, token}
}

func (service *UserService) Create(ctx context.Context, opts *CreateOptions) (*UserResponse, error) {
	response, err := service.client.Create(ctx, usersAPIPathname, service.token, newAPICreateOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalUserResponse(response.Body)
}

func (service *UserService) List(ctx context.Context, opts *ListOptions) ([]*UserResponse, bool, error) {
	response, err := service.client.List(ctx, usersAPIPathname, service.token, newAPIListOptions(opts))
	if err != nil {
		return nil, false, err
	}
	userPageResponse, err := unmarshalUserPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	return userPageResponse.Resources, len(userPageResponse.Resources) >= userPageResponse.ItemsPerPage, nil
}

func (service *UserService) Find(ctx context.Context, opts *FindOptions) (*UserResponse, error) {
	response, err := service.client.Find(ctx, usersAPIPathname, service.token, newAPIFindOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalUserResponse(response.Body)
}

func (service *UserService) Replace(ctx context.Context, opts *ReplaceOptions) (*UserResponse, error) {
	response, err := service.client.Replace(ctx, usersAPIPathname, service.token, newAPIReplaceOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalUserResponse(response.Body)
}

func (service *UserService) Update(ctx context.Context, opts *UpdateOptions) (bool, error) {
	_, err := service.client.Update(ctx, usersAPIPathname, service.token, newAPIUpdateOptions(opts))
	return err == nil, err
}

func (service *UserService) Delete(ctx context.Context, opts *DeleteOptions) (bool, error) {
	_, err := service.client.Delete(ctx, usersAPIPathname, service.token, newAPIDeleteOptions(opts))
	if err != nil {
		return false, err
	}
	return true, nil
}
