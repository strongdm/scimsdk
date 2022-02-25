package sdmscim

import "context"

type GroupService struct {
	token string
}

const GROUPS_API_PATHNAME = "Groups"

func newGroupService(token string) *GroupService {
	return &GroupService{token: token}
}

func (service *GroupService) create(ctx context.Context, opts *serviceCreateOptions) (*Group, error) {
	response, err := apiCreate(ctx, GROUPS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, err
	}
	unmarshedResponse, err := unmarshalGroupResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(unmarshedResponse), nil
}

func (service *GroupService) listIterator(ctx context.Context, opts *serviceListOptions) *GroupsIterator {
	listFunc := func(opts *serviceListOptions) ([]*Group, bool, error) {
		return service.list(ctx, opts)
	}
	return newGroupsIterator(listFunc, opts)
}

func (service *GroupService) list(ctx context.Context, opts *serviceListOptions) ([]*Group, bool, error) {
	response, err := apiList(ctx, GROUPS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, false, err
	}
	groupPageResponse, err := unmarshalGroupPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	groups := convertGroupResponseListToPorcelain(groupPageResponse.Resources)
	pageSize := defaultAPIPageSize
	if opts.PageSize != 0 {
		pageSize = opts.PageSize
	}
	return groups, len(groups) >= pageSize, nil
}

func (service *GroupService) find(ctx context.Context, opts *serviceFindOptions) (*Group, error) {
	response, err := apiFind(ctx, GROUPS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, err
	}
	unmarshedResponse, err := unmarshalGroupResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(unmarshedResponse), nil
}

// TODO: create method `replace`
func (service *GroupService) replace(ctx context.Context, opts *serviceReplaceOptions) (*Group, error) {
	response, err := apiReplace(ctx, GROUPS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, err
	}
	unmarshedResponse, err := unmarshalGroupResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(unmarshedResponse), nil
}

// TODO: create method `update`
func (service *GroupService) update(ctx context.Context, opts *serviceUpdateOptions) (bool, error) {
	_, err := apiUpdate(ctx, GROUPS_API_PATHNAME, service.token, opts)
	return err == nil, err
}

func (service *GroupService) delete(ctx context.Context, opts *serviceDeleteOptions) (bool, error) {
	_, err := apiDelete(ctx, GROUPS_API_PATHNAME, service.token, opts)
	return err == nil, err
}
