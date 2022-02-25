package sdmscim

import (
	"context"
)

type GroupModule struct {
	client  *Client
	service *GroupService
}

func (module *GroupModule) Create(ctx context.Context, group CreateGroupBody) (*Group, error) {
	body := convertPorcelainToCreateGroupRequest(&group)
	opts := newServiceCreateOptions(body, module.client.Options.APIUrl)
	return module.service.create(ctx, opts)
}

func (module *GroupModule) List(ctx context.Context, paginationOptions *PaginationOptions) *GroupsIterator {
	opts := newServiceListOptions(paginationOptions, module.client.Options.APIUrl)
	return module.service.listIterator(ctx, opts)
}

func (module *GroupModule) Find(ctx context.Context, id string) (*Group, error) {
	opts := newServiceFindOptions(id, module.client.Options.APIUrl)
	return module.service.find(ctx, opts)
}

func (module *GroupModule) Replace(ctx context.Context, id string, group ReplaceGroupBody) (*Group, error) {
	body := convertPorcelainToReplaceGroupRequest(&group)
	opts := newServiceReplaceOptions(id, body, module.client.Options.APIUrl)
	return module.service.replace(ctx, opts)
}

func (module *GroupModule) UpdateAddMembers(ctx context.Context, id string, members []UpdateGroupMemberBody) (bool, error) {
	body := convertPorcelainToUpdateGroupAddMembers(members)
	opts := newServiceUpdateOptions(id, body, module.client.Options.APIUrl)
	return module.service.update(ctx, opts)
}

func (module *GroupModule) UpdateReplaceName(ctx context.Context, id string, displayName string) {

}

func (module *GroupModule) UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error) {
	body := convertPorcelainToUpdateGroupRemoveMember(memberID)
	opts := newServiceUpdateOptions(id, body, module.client.Options.APIUrl)
	return module.service.update(ctx, opts)
}

func (module *GroupModule) Delete(ctx context.Context, id string) (bool, error) {
	opts := newServiceDeleteOptions(id, module.client.Options.APIUrl)
	return module.service.delete(ctx, opts)
}
