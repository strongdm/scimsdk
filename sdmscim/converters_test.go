package sdmscim

import (
	"sdmscim/sdmscim/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_USER_DISPLAY_NAME = "test user name"
)

func TestConvertUserResponseDTOToPorcelain(t *testing.T) {
	response := getUserResponseDTO()
	users := convertUserResponseListToPorcelain(response.Resources)
	firstUser := users[0]
	firstResponseUser := response.Resources[0]

	assert.Equal(t, firstResponseUser.ID, firstUser.ID)
	assert.Equal(t, firstResponseUser.DisplayName, firstUser.DisplayName)
	assert.Equal(t, firstResponseUser.Active, firstUser.Active)
	assert.Equal(t, firstResponseUser.ID, firstUser.ID)
	assert.Equal(t, firstResponseUser.ID, firstUser.ID)
	firstEmail := firstUser.Emails[0]
	firstResponseEmail := firstResponseUser.Emails[0]
	assert.Equal(t, firstEmail.Value, firstResponseEmail.Value)
	secondEmail := firstUser.Emails[1]
	secondResponseEmail := firstResponseUser.Emails[1]
	assert.Equal(t, secondEmail.Value, secondResponseEmail.Value)
	assert.Equal(t, firstResponseUser.UserName, firstUser.UserName)
	assert.Equal(t, firstResponseUser.UserType, firstUser.UserType)
}

func getUserResponseDTO() *api.APIUserPageResponseDTO {
	return &api.APIUserPageResponseDTO{
		Resources: []api.APIUserResponseDTO{
			{
				ID:          "xxx",
				Active:      true,
				DisplayName: "test user name",
				Emails: []api.APIUserEmailResponseDTO{
					{
						Primary: true,
						Value:   "username@email.com",
					},
					{
						Primary: false,
						Value:   "username2@email.com",
					},
				},
				Name: api.APIUserNameResponseDTO{
					FamilyName: "name",
					Formatted:  "test user name",
					GivenName:  "user",
				},
				Groups:   []interface{}{},
				Schemas:  []string{"test_schema"},
				UserName: "testuser",
				UserType: "account",
			},
		},
	}
}
