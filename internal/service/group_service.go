package service

import (
	"context"

	"github.com/strongdm/scimsdk/internal/api"
)

type GroupService interface {
	Create(ctx context.Context, opts *CreateOptions) (*GroupResponse, error)
	List(ctx context.Context, opts *ListOptions) ([]*GroupResponse, bool, error)
	Find(ctx context.Context, opts *FindOptions) (*GroupResponse, error)
	Replace(ctx context.Context, opts *ReplaceOptions) (*GroupResponse, error)
	Update(ctx context.Context, opts *UpdateOptions) (bool, error)
	Delete(ctx context.Context, opts *DeleteOptions) (bool, error)
}

type groupServiceImpl struct {
	client api.API
	token  string
}

const groupsAPIPathname = "Groups"

func NewGroupService(api api.API, token string) GroupService {
	return &groupServiceImpl{api, token}
}

func (service *groupServiceImpl) Create(ctx context.Context, opts *CreateOptions) (*GroupResponse, error) {
	response, err := service.client.Create(ctx, groupsAPIPathname, service.token, newAPICreateOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *groupServiceImpl) List(ctx context.Context, opts *ListOptions) ([]*GroupResponse, bool, error) {
	response, err := service.client.List(ctx, groupsAPIPathname, service.token, newAPIListOptions(opts))
	if err != nil {
		return nil, false, err
	}
	groupPageResponse, err := unmarshalGroupPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	return groupPageResponse.Resources, len(groupPageResponse.Resources) >= groupPageResponse.ItemsPerPage, nil
}

func (service *groupServiceImpl) Find(ctx context.Context, opts *FindOptions) (*GroupResponse, error) {
	response, err := service.client.Find(ctx, groupsAPIPathname, service.token, newAPIFindOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *groupServiceImpl) Replace(ctx context.Context, opts *ReplaceOptions) (*GroupResponse, error) {
	response, err := service.client.Replace(ctx, groupsAPIPathname, service.token, newAPIReplaceOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *groupServiceImpl) Update(ctx context.Context, opts *UpdateOptions) (bool, error) {
	_, err := service.client.Update(ctx, groupsAPIPathname, service.token, newAPIUpdateOptions(opts))
	return err == nil, err
}

func (service *groupServiceImpl) Delete(ctx context.Context, opts *DeleteOptions) (bool, error) {
	_, err := service.client.Delete(ctx, groupsAPIPathname, service.token, newAPIDeleteOptions(opts))
	return err == nil, err
}
