package sdmscim

import (
	"context"
)

type UserService struct {
	token string
	ctx   context.Context
}

func newUserService(token string, ctx context.Context) *UserService {
	return &UserService{token, ctx}
}

func (service *UserService) list(opts *ModuleListOptions) (users []*User, haveNextPage bool, err error) {
	apiOpts := moduleListOptionsToAPIOptions(opts)
	response, err := apiList(service.token, "Users", apiOpts, service.ctx)
	if err != nil {
		return nil, false, err
	}
	unmarshedResponse, err := unmarshalUserPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	users = convertUserResponseListToPorcelain(unmarshedResponse.Resources)
	pageSize := defaultAPIPageSize
	if opts.PageSize != 0 {
		pageSize = opts.PageSize
	}
	return users, len(users) >= pageSize, nil
}
