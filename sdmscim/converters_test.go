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
	t.Run("should convert user page response dto to user response porcelain list", func(t *testing.T) {
		response := getUserListResponseDTO()
		users := convertUserResponseListToPorcelain(response.Resources)
		firstUser := users[0]
		firstResponseUser := response.Resources[0]
		firstEmail := firstUser.Emails[0]
		firstResponseEmail := firstResponseUser.Emails[0]
		secondEmail := firstUser.Emails[1]
		secondResponseEmail := firstResponseUser.Emails[1]

		assert.Equal(t, firstResponseUser.ID, firstUser.ID)
		assert.Equal(t, firstResponseUser.DisplayName, firstUser.DisplayName)
		assert.Equal(t, firstResponseUser.Active, firstUser.Active)
		assert.Equal(t, firstResponseUser.ID, firstUser.ID)
		assert.Equal(t, firstResponseUser.ID, firstUser.ID)
		assert.Equal(t, firstResponseEmail.Value, firstEmail.Value)
		assert.Equal(t, secondResponseEmail.Value, secondEmail.Value)
		assert.Equal(t, firstResponseUser.UserName, firstUser.UserName)
		assert.Equal(t, firstResponseUser.UserType, firstUser.UserType)
	})
}

func getUserListResponseDTO() *api.APIUserPageResponse {
	return &api.APIUserPageResponse{
		Resources: []api.APIUserResponse{
			{
				ID:          "xxx",
				Active:      true,
				DisplayName: "test user name",
				Emails: []api.APIUserEmailResponse{
					{
						Primary: true,
						Value:   "username@email.com",
					},
					{
						Primary: false,
						Value:   "username2@email.com",
					},
				},
				Name: api.APIUserNameResponse{
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
