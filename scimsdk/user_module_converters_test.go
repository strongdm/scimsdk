package scimsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockUserID = "xxx"

func TestConvertUserToAndFromPorcelain(t *testing.T) {
	t.Run("should convert a replace user body to api body when passing a valid replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		apiBody, err := convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.Nil(t, err)
		assert.Equal(t, mockUserID, apiBody.ID)
		assert.Equal(t, body.UserName, apiBody.UserName)
		assert.Equal(t, body.GivenName, apiBody.Name.GivenName)
		assert.Equal(t, body.FamilyName, apiBody.Name.FamilyName)
		assert.Equal(t, body.Active, apiBody.Active)
	})

	t.Run("should return an error when passing an empty id to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		_, err := convertPorcelainToReplaceUserRequest("", body)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the user id")
	})

	t.Run("should return an error when passing an empty userName to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		body.UserName = ""
		_, err := convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "must pass the user email")
	})

	t.Run("should return an error when passing an empty givenName to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		body.GivenName = ""
		_, err := convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the user first name")
	})

	t.Run("should return an error when passing an empty familyName to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		body.FamilyName = ""
		_, err := convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the user last name")
	})

	t.Run("should convert a create user body to api body when passing a valid create user body", func(t *testing.T) {
		body := getValidCreateUser()
		apiBody, err := convertPorcelainToCreateUserRequest(body)

		assert.Nil(t, err)
		assert.Equal(t, body.UserName, apiBody.UserName)
		assert.Equal(t, body.GivenName, apiBody.Name.GivenName)
		assert.Equal(t, body.FamilyName, apiBody.Name.FamilyName)
		assert.Equal(t, body.Active, apiBody.Active)
	})

	t.Run("should return an error when passing an empty userName to api create user body", func(t *testing.T) {
		body := getValidCreateUser()
		body.UserName = ""
		_, err := convertPorcelainToCreateUserRequest(body)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the user email")
	})

	t.Run("should return an error when passing an empty givenName to api create user body", func(t *testing.T) {
		body := getValidCreateUser()
		body.GivenName = ""
		_, err := convertPorcelainToCreateUserRequest(body)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "must pass the user first name")
	})

	t.Run("should return an error when passing an empty familyName to api create user body", func(t *testing.T) {
		body := getValidCreateUser()
		body.FamilyName = ""
		_, err := convertPorcelainToCreateUserRequest(body)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "must pass the user last name")
	})
}

func getValidCreateUser() *CreateUser {
	return &CreateUser{
		UserName:   "xxx",
		GivenName:  "yyy",
		FamilyName: "zzz",
		Active:     true,
	}
}

func getValidReplaceUser() *ReplaceUser {
	return &ReplaceUser{
		UserName:   "xxx",
		GivenName:  "yyy",
		FamilyName: "zzz",
		Active:     true,
	}
}
