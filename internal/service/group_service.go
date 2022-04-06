package service

import (
	"context"

	"github.com/strongdm/scimsdk/internal/api"
)

type IGroupService interface {
	Create(ctx context.Context, opts *CreateOptions) (*GroupResponse, error)
	List(ctx context.Context, opts *ListOptions) ([]*GroupResponse, bool, error)
	Find(ctx context.Context, opts *FindOptions) (*GroupResponse, error)
	Replace(ctx context.Context, opts *ReplaceOptions) (*GroupResponse, error)
	Update(ctx context.Context, opts *UpdateOptions) (bool, error)
	Delete(ctx context.Context, opts *DeleteOptions) (bool, error)
}

type GroupService struct {
	client api.IAPI
	token  string
}

const groupsAPIPathname = "Groups"

func NewGroupService(api api.IAPI, token string) IGroupService {
	return &GroupService{api, token}
}

func (service *GroupService) Create(ctx context.Context, opts *CreateOptions) (*GroupResponse, error) {
	response, err := service.client.Create(ctx, groupsAPIPathname, service.token, newAPICreateOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *GroupService) List(ctx context.Context, opts *ListOptions) ([]*GroupResponse, bool, error) {
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

func (service *GroupService) Find(ctx context.Context, opts *FindOptions) (*GroupResponse, error) {
	response, err := service.client.Find(ctx, groupsAPIPathname, service.token, newAPIFindOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *GroupService) Replace(ctx context.Context, opts *ReplaceOptions) (*GroupResponse, error) {
	response, err := service.client.Replace(ctx, groupsAPIPathname, service.token, newAPIReplaceOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *GroupService) Update(ctx context.Context, opts *UpdateOptions) (bool, error) {
	_, err := service.client.Update(ctx, groupsAPIPathname, service.token, newAPIUpdateOptions(opts))
	return err == nil, err
}

func (service *GroupService) Delete(ctx context.Context, opts *DeleteOptions) (bool, error) {
	_, err := service.client.Delete(ctx, groupsAPIPathname, service.token, newAPIDeleteOptions(opts))
	return err == nil, err
}
