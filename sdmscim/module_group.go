package sdmscim

type GroupModule struct {
	client *Client
}

func (service GroupModule) List() *GroupsIterator {
	api := newGroupService(service.client.adminToken)
	iterator := newGroupsIterator(api.list)
	return iterator
}
