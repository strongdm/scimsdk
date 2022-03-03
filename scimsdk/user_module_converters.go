package scimsdk

import (
	"log"
	"scimsdk/internal/service"
)

func convertUserResponseListToPorcelain(response []service.UserResponse) []*User {
	users := []*User{}
	for _, item := range response {
		users = append(users, convertUserResponseToPorcelain(&item))
	}
	return users
}

func convertUserResponseToPorcelain(response *service.UserResponse) *User {
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

func convertUserNameResponseToPorcelain(response service.UserNameResponse) *UserName {
	return &UserName{
		Formatted:  response.Formatted,
		FamilyName: response.FamilyName,
		GivenName:  response.GivenName,
	}
}

func convertUserEmailResponseListToPorcelain(response []service.UserEmailResponse) []UserEmail {
	emails := []UserEmail{}
	for _, userEmail := range response {
		emails = append(emails, convertUserEmailResponseToPorcelain(&userEmail))
	}
	return emails
}

func convertUserEmailResponseToPorcelain(response *service.UserEmailResponse) UserEmail {
	return UserEmail{
		Primary: response.Primary,
		Value:   response.Value,
	}
}

func convertPorcelainToCreateUserRequest(user *CreateUser) *service.CreateUserRequest {
	if user.UserName == "" {
		log.Fatal("You must pass the user email in UserName field.")
	} else if user.GivenName == "" {
		log.Fatal("You must pass the user first name in GivenName field.")
	} else if user.FamilyName == "" {
		log.Fatal("You must pass the user last name in FamilyName field.")
	}
	return &service.CreateUserRequest{
		Schemas:  []string{defaultUserSchema},
		UserName: user.UserName,
		Name:     service.UserNameRequest{GivenName: user.GivenName, FamilyName: user.FamilyName},
		Active:   user.Active,
	}
}

func convertPorcelainToReplaceUserRequest(id string, user *ReplaceUser) *service.ReplaceUserRequest {
	if id == "" {
		log.Fatal("You must pass the user id.")
	} else if user.UserName == "" {
		log.Fatal("You must pass the user email in UserName field.")
	} else if user.GivenName == "" {
		log.Fatal("You must pass the user first name in GivenName field.")
	} else if user.FamilyName == "" {
		log.Fatal("You must pass the user last name in FamilyName field.")
	}
	return &service.ReplaceUserRequest{
		ID:       id,
		Schemas:  []string{defaultUserSchema},
		UserName: user.UserName,
		Name:     service.UserNameRequest{GivenName: user.GivenName, FamilyName: user.FamilyName},
		Active:   user.Active,
	}
}

func convertPorcelainToUpdateUserRequest(body UpdateUser) *service.UpdateUserRequest {
	return &service.UpdateUserRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []service.UpdateUserOperationRequest{
			{
				OP:    "replace",
				Value: service.UpdateUserOperationValueRequest(body),
			},
		},
	}
}