package sdmscim

import (
	"context"
	"sdmscim/sdmscim/api"
)

type UserService struct {
	token string
	ctx   context.Context
}

func newUserService(token string, ctx context.Context) *UserService {
	return &UserService{token, ctx}
}

// TODO: Add opts as params (opts == {paginationLimit: 0, filter: ""})
func (service *UserService) list(offset int) (users []*User, haveNextPage bool, err error) {
	response, err := api.List(service.token, "Users", offset, service.ctx)
	if err != nil {
		return nil, false, err
	}
	unmarshedResponse, err := unmarshalUserPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	users = convertUserResponseListToPorcelain(unmarshedResponse.Resources)
	return users, len(users) >= api.DEFAULT_USERS_PAGE_LIMIT, nil
}
