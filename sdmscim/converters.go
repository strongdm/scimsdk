package sdmscim

import (
	"encoding/json"
	"io"
	"log"
)

const defaultUserSchema = "urn:ietf:params:scim:schemas:core:2.0:User"

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

func unmarshalUserResponse(body io.ReadCloser) (*apiUserResponse, error) {
	unmarshedResponse := &apiUserResponse{}
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

func convertPorcelainToCreateUserRequest(user *CreateUserBody) *apiCreateUserRequest {
	if user.UserName == "" {
		log.Fatal("You must pass the user email in UserName field.")
	} else if user.GivenName == "" {
		log.Fatal("You must pass the user first name in GivenName field.")
	} else if user.FamilyName == "" {
		log.Fatal("You must pass the user last name in FamilyName field.")
	}
	return &apiCreateUserRequest{
		Schemas:  []string{defaultUserSchema},
		UserName: user.UserName,
		Name:     apiUserNameRequest{user.GivenName, user.FamilyName},
		Active:   user.Active,
	}
}

func convertPorcelainToReplaceUserRequest(user *ReplaceUserBody) *apiReplaceUserRequest {
	if user.UserName == "" {
		log.Fatal("You must pass the user email in UserName field.")
	} else if user.GivenName == "" {
		log.Fatal("You must pass the user first name in GivenName field.")
	} else if user.FamilyName == "" {
		log.Fatal("You must pass the user last name in FamilyName field.")
	}
	return &apiReplaceUserRequest{
		ID:       user.ID,
		Schemas:  []string{defaultUserSchema},
		UserName: user.UserName,
		Name:     apiUserNameRequest{user.GivenName, user.FamilyName},
		Active:   user.Active,
	}
}

func convertPorcelainToUpdateUserRequest(active bool) *apiUpdateUserRequest {
	return &apiUpdateUserRequest{
		Schemas: []string{defaultUserSchema},
		Operations: []apiUpdateOperationRequest{
			{
				OP: "replace",
				Value: apiUpdateOperationValueRequest{
					Active: active,
				},
			},
		},
	}
}
