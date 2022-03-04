package scimsdk

import (
	"context"

	"github.com/strongdm/scimsdk/internal/service"
)

type GroupModule struct {
	client  *Client
	service *service.GroupService
}

func (module *GroupModule) Create(ctx context.Context, group CreateGroupBody) (*Group, error) {
	body, err := convertPorcelainToCreateGroupRequest(&group)
	if err != nil {
		return nil, err
	}
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
	opts, err := newServiceFindOptions(id, module.client.GetProvidedURL())
	if err != nil {
		return nil, err
	}
	response, err := module.service.Find(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModule) Replace(ctx context.Context, id string, group ReplaceGroupBody) (*Group, error) {
	body, err := convertPorcelainToReplaceGroupRequest(&group)
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
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModule) UpdateAddMembers(ctx context.Context, id string, members []GroupMember) (bool, error) {
	body, err := convertPorcelainToUpdateGroupAddMembersRequest(members)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) UpdateReplaceMembers(ctx context.Context, id string, members []GroupMember) (bool, error) {
	body, err := convertPorcelainToUpdateGroupReplaceMembersRequest(members)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) UpdateReplaceName(ctx context.Context, id string, replaceName UpdateGroupReplaceName) (bool, error) {
	body, err := convertPorcelainToUpdateGroupNameRequest(replaceName)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error) {
	body, err := convertPorcelainToUpdateGroupRemoveMemberRequest(memberID)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModule) Delete(ctx context.Context, id string) (bool, error) {
	opts, err := newServiceDeleteOptions(id, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Delete(ctx, opts)
}

func (module *GroupModule) iteratorMiddleware(ctx context.Context) func(opts *PaginationOptions) ([]*Group, bool, error) {
	return func(opts *PaginationOptions) ([]*Group, bool, error) {
		listOpts, err := newServiceListOptions(opts, module.client.GetProvidedURL())
		if err != nil {
			return nil, false, err
		}
		response, haveNextPage, err := module.service.List(ctx, listOpts)
		if err != nil {
			return nil, false, err
		}
		groups := convertGroupResponseListToPorcelain(response)
		return groups, haveNextPage, nil
	}
}
