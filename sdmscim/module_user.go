package sdmscim

type UserModule struct {
	client *Client
}

func (service UserModule) List() *UsersIterator {
	// TODO: test more to see if it's necessary to use factories instead constructing directly
	api := newUserService(service.client.adminToken)
	iterator := newUsersIterator(api.list)
	return iterator
}
