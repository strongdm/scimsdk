package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/strongdm/scimsdk/models"
)

const mockUserID = "xxx"

func TestConvertUserToAndFromPorcelain(t *testing.T) {
	t.Run("should convert a replace user body to api body when passing a valid replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		apiBody, err := convertPorcelainToReplaceUserRequest(mockUserID, body)
		assertT := assert.New(t)

		assertT.Nil(err)
		assertT.Equal(mockUserID, apiBody.ID)
		assertT.Equal(body.UserName, apiBody.UserName)
		assertT.Equal(body.GivenName, apiBody.Name.GivenName)
		assertT.Equal(body.FamilyName, apiBody.Name.FamilyName)
		assertT.Equal(body.Active, apiBody.Active)
	})

	t.Run("should return an error when passing an empty id to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		_, err := convertPorcelainToReplaceUserRequest("", body)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the user id")
	})

	t.Run("should return an error when passing an empty userName to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		body.UserName = ""
		_, err := convertPorcelainToReplaceUserRequest(mockUserID, body)
		assertT := assert.New(t)

		assertT.NotNil(err)
		assertT.Contains(err.Error(), "must pass the user email")
	})

	t.Run("should return an error when passing an empty givenName to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		body.GivenName = ""
		_, err := convertPorcelainToReplaceUserRequest(mockUserID, body)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the user first name")
	})

	t.Run("should return an error when passing an empty familyName to api replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		body.FamilyName = ""
		_, err := convertPorcelainToReplaceUserRequest(mockUserID, body)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the user last name")
	})

	t.Run("should convert a create user body to api body when passing a valid create user body", func(t *testing.T) {
		body := getValidCreateUser()
		apiBody, err := convertPorcelainToCreateUserRequest(body)
		assertT := assert.New(t)

		assertT.Nil(err)
		assertT.Equal(body.UserName, apiBody.UserName)
		assertT.Equal(body.GivenName, apiBody.Name.GivenName)
		assertT.Equal(body.FamilyName, apiBody.Name.FamilyName)
		assertT.Equal(body.Active, apiBody.Active)
	})

	t.Run("should return an error when passing an empty userName to api create user body", func(t *testing.T) {
		body := getValidCreateUser()
		body.UserName = ""
		_, err := convertPorcelainToCreateUserRequest(body)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the user email")
	})

	t.Run("should return an error when passing an empty givenName to api create user body", func(t *testing.T) {
		body := getValidCreateUser()
		body.GivenName = ""
		_, err := convertPorcelainToCreateUserRequest(body)
		assertT := assert.New(t)

		assertT.NotNil(err)
		assertT.Contains(err.Error(), "must pass the user first name")
	})

	t.Run("should return an error when passing an empty familyName to api create user body", func(t *testing.T) {
		body := getValidCreateUser()
		body.FamilyName = ""
		_, err := convertPorcelainToCreateUserRequest(body)
		assertT := assert.New(t)

		assertT.NotNil(err)
		assertT.Contains(err.Error(), "must pass the user last name")
	})
}

func getValidCreateUser() *models.CreateUser {
	return &models.CreateUser{
		UserName:   "xxx",
		GivenName:  "yyy",
		FamilyName: "zzz",
		Active:     true,
	}
}

func getValidReplaceUser() *models.ReplaceUser {
	return &models.ReplaceUser{
		UserName:   "xxx",
		GivenName:  "yyy",
		FamilyName: "zzz",
		Active:     true,
	}
}
