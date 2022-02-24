package sdmscim

import (
	"context"
)

type UserService struct {
	token string
}

const USERS_API_PATHNAME = "Users"

func newUserService(token string) *UserService {
	return &UserService{token: token}
}

func (service *UserService) create(ctx context.Context, opts *serviceCreateOptions) (*User, error) {
	response, err := apiCreate(ctx, USERS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, err
	}
	unmarshedResponse, err := unmarshalUserResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(unmarshedResponse), nil
}

func (service *UserService) listIterator(ctx context.Context, opts *serviceListOptions) *UsersIterator {
	listFunc := func(opts *serviceListOptions) ([]*User, bool, error) {
		return service.list(ctx, opts)
	}
	return newUsersIterator(listFunc, opts)
}

func (service *UserService) list(ctx context.Context, opts *serviceListOptions) ([]*User, bool, error) {
	response, err := apiList(ctx, USERS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, false, err
	}
	unmarshedResponse, err := unmarshalUserPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	users := convertUserResponseListToPorcelain(unmarshedResponse.Resources)
	pageSize := defaultAPIPageSize
	if opts.PageSize != 0 {
		pageSize = opts.PageSize
	}
	return users, len(users) >= pageSize, nil
}

func (service *UserService) find(ctx context.Context, opts *serviceFindOptions) (*User, error) {
	response, err := apiFind(ctx, USERS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, err
	}
	unmarshedResponse, err := unmarshalUserResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(unmarshedResponse), nil
}

func (service *UserService) replace(ctx context.Context, opts *serviceReplaceOptions) (*User, error) {
	response, err := apiReplace(ctx, USERS_API_PATHNAME, service.token, opts)
	if err != nil {
		return nil, err
	}
	unmarshedResponse, err := unmarshalUserResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return convertUserResponseToPorcelain(unmarshedResponse), nil
}

func (service *UserService) delete(ctx context.Context, opts *serviceDeleteOptions) (bool, error) {
	_, err := apiDelete(ctx, USERS_API_PATHNAME, service.token, opts)
	if err != nil {
		return false, err
	}
	return true, nil
}
