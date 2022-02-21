package sdmscim

import (
	"sdmscim/sdmscim/api"
)

type UserService struct {
	token string
}

func newUserService(token string) *UserService {
	return &UserService{token}
}

// TODO: Add opts as params (opts == {paginationLimit: 0, filter: ""})
func (a *UserService) list(offset int) (users []*User, haveNextPage bool, err error) {
	response, err := api.List(a.token, "Users", offset)
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
