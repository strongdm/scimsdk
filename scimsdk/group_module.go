package scimsdk

import (
	"context"

	"github.com/strongdm/scimsdk/internal/service"
)

type GroupModule interface {
	Create(ctx context.Context, group CreateGroupBody) (*Group, error)
	List(ctx context.Context, paginationOptions *PaginationOptions) GroupIterator
	Find(ctx context.Context, id string) (*Group, error)
	Replace(ctx context.Context, id string, group ReplaceGroupBody) (*Group, error)
	UpdateAddMembers(ctx context.Context, id string, members []GroupMember) (bool, error)
	UpdateReplaceMembers(ctx context.Context, id string, members []GroupMember) (bool, error)
	UpdateReplaceName(ctx context.Context, id string, replaceName UpdateGroupReplaceName) (bool, error)
	UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type groupModuleImpl struct {
	client  Client
	service service.GroupService
}

func NewGroupModule(client Client, service service.GroupService) GroupModule {
	return &groupModuleImpl{client, service}
}

func (module *groupModuleImpl) Create(ctx context.Context, group CreateGroupBody) (*Group, error) {
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

func (module *groupModuleImpl) List(ctx context.Context, paginationOptions *PaginationOptions) GroupIterator {
	return newGroupsIterator(module.iteratorMiddleware(ctx), paginationOptions)
}

func (module *groupModuleImpl) Find(ctx context.Context, id string) (*Group, error) {
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

func (module *groupModuleImpl) Replace(ctx context.Context, id string, group ReplaceGroupBody) (*Group, error) {
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

func (module *groupModuleImpl) UpdateAddMembers(ctx context.Context, id string, members []GroupMember) (bool, error) {
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

func (module *groupModuleImpl) UpdateReplaceMembers(ctx context.Context, id string, members []GroupMember) (bool, error) {
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

func (module *groupModuleImpl) UpdateReplaceName(ctx context.Context, id string, replaceName UpdateGroupReplaceName) (bool, error) {
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

func (module *groupModuleImpl) UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error) {
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

func (module *groupModuleImpl) Delete(ctx context.Context, id string) (bool, error) {
	opts, err := newServiceDeleteOptions(id, module.client.GetProvidedURL())
	if err != nil {
		return false, err
	}
	return module.service.Delete(ctx, opts)
}

func (module *groupModuleImpl) iteratorMiddleware(ctx context.Context) listGroupsOperationFunc {
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
