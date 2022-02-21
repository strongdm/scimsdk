package sdmscim

import (
	"encoding/json"
	"io"
	"sdmscim/sdmscim/api"
)

// ----------------------------------

func unmarshalUserPageResponse(body io.ReadCloser) (*api.APIUserPageResponse, error) {
	unmarshedResponse := &api.APIUserPageResponse{}
	buff, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buff, &unmarshedResponse)
	if err != nil {
		return nil, err
	}
	return unmarshedResponse, nil
}

func unmarshalGroupPageResponse(body io.ReadCloser) (*api.APIGroupPageResponse, error) {
	unmarshedResponse := &api.APIGroupPageResponse{}
	buff, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buff, &unmarshedResponse)
	if err != nil {
		return nil, err
	}
	return unmarshedResponse, nil
}

func convertUserResponseListToPorcelain(response []api.APIUserResponse) []*User {
	users := []*User{}
	for _, item := range response {
		users = append(users, convertUserResponseToPorcelain(&item))
	}
	return users
}

func convertUserResponseToPorcelain(response *api.APIUserResponse) *User {
	return &User{
		ID:          response.ID,
		Active:      response.Active,
		DisplayName: response.DisplayName,
		Emails:      convertUserEmailResponseListToPorcelain(response.Emails),
		Groups:      response.Groups,
		Name:        convertUserNameResponseToPorcelain(response.Name),
		UserName:    response.UserName,
		UserType:    response.UserType,
	}
}

func convertUserNameResponseToPorcelain(response api.APIUserNameResponse) *UserName {
	return &UserName{
		Formatted:  response.Formatted,
		FamilyName: response.FamilyName,
		GivenName:  response.GivenName,
	}
}

func convertUserEmailResponseListToPorcelain(response []api.APIUserEmailResponse) []UserEmail {
	emails := []UserEmail{}
	for _, userEmail := range response {
		emails = append(emails, convertUserEmailResponseToPorcelain(&userEmail))
	}
	return emails
}

func convertUserEmailResponseToPorcelain(response *api.APIUserEmailResponse) UserEmail {
	return UserEmail{
		Primary: response.Primary,
		Value:   response.Value,
	}
}

func convertGroupResponseListToPorcelain(response []api.APIGroupResponse) []*Group {
	groups := []*Group{}
	for _, item := range response {
		groups = append(groups, convertGroupResponseToPorcelain(&item))
	}
	return groups
}

func convertGroupResponseToPorcelain(response *api.APIGroupResponse) *Group {
	return &Group{
		ID:          response.ID,
		DisplayName: response.DisplayName,
		Members:     response.Members,
		Meta:        convertGroupMetaResponseToPorcelain(&response.Meta),
	}
}

func convertGroupMetaResponseToPorcelain(response *api.APIGroupMetaResponse) *GroupMeta {
	return &GroupMeta{
		ResourceType: response.ResourceType,
		Location:     response.Location,
	}
}
