package sdmscim

import "context"

type GroupModule struct {
	client *Client
}

func (module GroupModule) List(ctx context.Context) *GroupsIterator {
	service := newGroupService(module.client.adminToken, ctx)
	iterator := newGroupsIterator(service.list)
	return iterator
}
