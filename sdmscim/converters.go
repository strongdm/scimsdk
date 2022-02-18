package sdmscim

import (
	"encoding/json"
	"io"
	"sdmscim/sdmscim/api"
)

// ----------------------------------

func unmarshalUserPageResponse(body io.ReadCloser) (*api.APIUserPageResponseDTO, error) {
	unmarshedResponse := &api.APIUserPageResponseDTO{}
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

func unmarshalGroupPageResponse(body io.ReadCloser) (*api.APIGroupPageResponseDTO, error) {
	unmarshedResponse := &api.APIGroupPageResponseDTO{}
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

func convertUserResponseListToPorcelain(response []api.APIUserResponseDTO) []*User {
	users := []*User{}
	for _, item := range response {
		users = append(users, convertUserResponseToPorcelain(&item))
	}
	return users
}

func convertUserResponseToPorcelain(response *api.APIUserResponseDTO) *User {
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

func convertUserNameResponseToPorcelain(response api.APIUserNameResponseDTO) *UserName {
	return &UserName{
		Formatted:  response.Formatted,
		FamilyName: response.FamilyName,
		GivenName:  response.GivenName,
	}
}

func convertUserEmailResponseListToPorcelain(response []api.APIUserEmailResponseDTO) []UserEmail {
	emails := []UserEmail{}
	for _, userEmail := range response {
		emails = append(emails, convertUserEmailResponseToPorcelain(&userEmail))
	}
	return emails
}

func convertUserEmailResponseToPorcelain(response *api.APIUserEmailResponseDTO) UserEmail {
	return UserEmail{
		Primary: response.Primary,
		Value:   response.Value,
	}
}

func convertGroupResponseListToPorcelain(response []api.APIGroupResponseDTO) []*Group {
	groups := []*Group{}
	for _, item := range response {
		groups = append(groups, convertGroupResponseToPorcelain(&item))
	}
	return groups
}

func convertGroupResponseToPorcelain(response *api.APIGroupResponseDTO) *Group {
	return &Group{
		ID:          response.ID,
		DisplayName: response.DisplayName,
		Members:     response.Members,
		Meta:        convertGroupMetaResponseToPorcelain(&response.Meta),
	}
}

func convertGroupMetaResponseToPorcelain(response *api.APIGroupMetaResponseDTO) *GroupMeta {
	return &GroupMeta{
		ResourceType: response.ResourceType,
		Location:     response.Location,
	}
}
