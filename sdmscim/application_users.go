package sdmscim

import (
	"sdmscim/sdmscim/api"
)

type UserApplication struct {
	token string
}

func newUserApplication(token string) *UserApplication {
	return &UserApplication{token}
}

func (a *UserApplication) list(offset int) (users []*User, haveNextPage bool, err error) {
	response, err := api.BaseList(a.token, "Users", offset)
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
