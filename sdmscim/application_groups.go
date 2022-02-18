package sdmscim

import "sdmscim/sdmscim/api"

type GroupApplication struct {
	token string
}

func newGroupApplication(token string) *GroupApplication {
	return &GroupApplication{token}
}

// TODO: Support filter param
func (a *GroupApplication) list(offset int) (groups []*Group, haveNextPage bool, err error) {
	response, err := api.BaseList(a.token, "Groups", offset)
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
