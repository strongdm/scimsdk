package sdmscim

import "sdmscim/sdmscim/api"

type GroupService struct {
	token string
}

func newGroupService(token string) *GroupService {
	return &GroupService{token}
}

// TODO: Support filter param
func (a *GroupService) list(offset int) (groups []*Group, haveNextPage bool, err error) {
	response, err := api.List(a.token, "Groups", offset)
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
