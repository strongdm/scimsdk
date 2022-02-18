package sdmscim

type UserService struct {
	client *Client
}

func (service UserService) List() *usersIterator {
	api := newUserApplication(service.client.adminToken)
	iterator := newUsersIterator(api.list)
	return iterator
}
