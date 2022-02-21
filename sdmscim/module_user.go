package sdmscim

import "context"

type UserModule struct {
	client *Client
}

func (module UserModule) List(ctx context.Context) *UsersIterator {
	service := newUserService(module.client.adminToken, ctx)
	iterator := newUsersIterator(service.list)
	return iterator
}
