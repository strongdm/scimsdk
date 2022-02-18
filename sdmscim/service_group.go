package sdmscim

type GroupService struct {
	client *Client
}

func (service GroupService) List() *groupsIterator {
	api := newGroupApplication(service.client.adminToken)
	iterator := newGroupsIterator(api.list)
	return iterator
}
