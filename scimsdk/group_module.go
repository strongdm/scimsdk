package scimsdk

import (
	"context"
	"scimsdk/internal/service"
)

type GroupModule struct {
	client  *Client
	service *service.GroupService
}

func (module *GroupModule) Create(ctx context.Context, group CreateGroupBody) (*Group, error) {
	body := convertPorcelainToCreateGroupRequest(&group)
	opts := newServiceCreateOptions(body, module.client.GetProvidedURL())
	response, err := module.service.Create(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModule) List(ctx context.Context, paginationOptions *PaginationOptions) *GroupsIterator {
	return newGroupsIterator(module.iteratorMiddleware(ctx), paginationOptions)
}

func (module *GroupModule) Find(ctx context.Context, id string) (*Group, error) {
	opts := newServiceFindOptions(id, module.client.GetProvidedURL())
	response, err := module.service.Find(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModule) Replace(ctx context.Context, id string, group ReplaceGroupBody) (*Group, error) {
	body := convertPorcelainToReplaceGroupRequest(&group)
	opts := newServiceReplaceOptions(id, body, module.client.GetProvidedURL())
	response, err := module.service.Replace(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModule) UpdateAddMembers(ctx context.Context, id string, members []GroupMember) (bool, error) {
	body := convertPorcelainToUpdateGroupAddMembersRequest(members)
	opts := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) UpdateReplaceMembers(ctx context.Context, id string, members []GroupMember) (bool, error) {
	body := convertPorcelainToUpdateGroupReplaceMembersRequest(members)
	opts := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) UpdateReplaceName(ctx context.Context, id string, replaceName UpdateGroupReplaceName) (bool, error) {
	body := convertPorcelainToUpdateGroupNameRequest(replaceName)
	opts := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error) {
	body := convertPorcelainToUpdateGroupRemoveMemberRequest(memberID)
	opts := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) Delete(ctx context.Context, id string) (bool, error) {
	opts := newServiceDeleteOptions(id, module.client.GetProvidedURL())
	return module.service.Delete(ctx, opts)
}

func (module *GroupModule) iteratorMiddleware(ctx context.Context) func(opts *PaginationOptions) ([]*Group, bool, error) {
	return func(opts *PaginationOptions) ([]*Group, bool, error) {
		listOpts := newServiceListOptions(opts, module.client.GetProvidedURL())
		response, haveNextPage, err := module.service.List(ctx, listOpts)
		if err != nil {
			return nil, false, err
		}
		groups := convertGroupResponseListToPorcelain(response)
		return groups, haveNextPage, nil
	}
}
