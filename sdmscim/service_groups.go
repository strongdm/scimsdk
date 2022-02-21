package sdmscim

import (
	"context"
	"sdmscim/sdmscim/api"
)

type GroupService struct {
	token string
	ctx   context.Context
}

func newGroupService(token string, ctx context.Context) *GroupService {
	return &GroupService{token, ctx}
}

func (service *GroupService) list(offset int) (groups []*Group, haveNextPage bool, err error) {
	response, err := api.List(service.token, "Groups", offset, service.ctx)
	if err != nil {
		return nil, false, err
	}
	unmarshedResponse, err := unmarshalGroupPageResponse(response.Body)
	if err != nil {
		return nil, false, err
	}
	groups = convertGroupResponseListToPorcelain(unmarshedResponse.Resources)
	return groups, len(groups) >= api.DEFAULT_USERS_PAGE_LIMIT, nil
}
