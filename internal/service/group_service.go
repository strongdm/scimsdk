package service

import (
	"context"
	"scimsdk/internal/api"
)

type GroupService struct {
	token string
}

const groupsAPIPathname = "Groups"

func NewGroupService(token string) *GroupService {
	return &GroupService{token: token}
}

func (service *GroupService) Create(ctx context.Context, opts *CreateOptions) (*GroupResponse, error) {
	response, err := api.Create(ctx, groupsAPIPathname, service.token, newAPICreateOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *GroupService) ListIteratorMiddleware(ctx context.Context) func(opts *ListOptions) ([]*GroupResponse, bool, error) {
	return func(opts *ListOptions) ([]*GroupResponse, bool, error) {
		return service.List(ctx, opts)
	}
}

func (service *GroupService) List(ctx context.Context, opts *ListOptions) ([]*GroupResponse, bool, error) {
	response, err := api.List(ctx, groupsAPIPathname, service.token, newAPIListOptions(opts))
	if err != nil {
		return nil, false, err
	}
	groupPageResponse, err := unmarshalGroupPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	pageSize := api.GetDefaultAPIPageSize()
	if opts.PageSize != 0 {
		pageSize = opts.PageSize
	}
	return groupPageResponse.Resources, len(groupPageResponse.Resources) >= pageSize, nil
}

func (service *GroupService) Find(ctx context.Context, opts *FindOptions) (*GroupResponse, error) {
	response, err := api.Find(ctx, groupsAPIPathname, service.token, newAPIFindOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *GroupService) Replace(ctx context.Context, opts *ReplaceOptions) (*GroupResponse, error) {
	response, err := api.Replace(ctx, groupsAPIPathname, service.token, newAPIReplaceOptions(opts))
	if err != nil {
		return nil, err
	}
	return unmarshalGroupResponse(response.Body)
}

func (service *GroupService) Update(ctx context.Context, opts *UpdateOptions) (bool, error) {
	_, err := api.Update(ctx, groupsAPIPathname, service.token, newAPIUpdateOptions(opts))
	return err == nil, err
}

func (service *GroupService) Delete(ctx context.Context, opts *DeleteOptions) (bool, error) {
	_, err := api.Delete(ctx, groupsAPIPathname, service.token, newAPIDeleteOptions(opts))
	return err == nil, err
}
