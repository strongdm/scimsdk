package sdmscim

import (
	"encoding/json"
	"io"
)

func unmarshalUserPageResponse(body io.ReadCloser) (*apiUserPageResponse, error) {
	unmarshedResponse := &apiUserPageResponse{}
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

func convertUserResponseListToPorcelain(response []apiUserResponse) []*User {
	users := []*User{}
	for _, item := range response {
		users = append(users, convertUserResponseToPorcelain(&item))
	}
	return users
}

func convertUserResponseToPorcelain(response *apiUserResponse) *User {
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

func convertUserNameResponseToPorcelain(response apiUserNameResponse) *UserName {
	return &UserName{
		Formatted:  response.Formatted,
		FamilyName: response.FamilyName,
		GivenName:  response.GivenName,
	}
}

func convertUserEmailResponseListToPorcelain(response []apiUserEmailResponse) []UserEmail {
	emails := []UserEmail{}
	for _, userEmail := range response {
		emails = append(emails, convertUserEmailResponseToPorcelain(&userEmail))
	}
	return emails
}

func convertUserEmailResponseToPorcelain(response *apiUserEmailResponse) UserEmail {
	return UserEmail{
		Primary: response.Primary,
		Value:   response.Value,
	}
}
