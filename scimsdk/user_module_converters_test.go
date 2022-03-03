package scimsdk

import (
	"log"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

const mockUserID = "xxx"

func TestConvertUserToAndFromPorcelain(t *testing.T) {
	t.Run("should convert a replace user body to api body when passing a valid replace user body", func(t *testing.T) {
		body := getValidReplaceUser()
		apiBody := convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.Equal(t, mockUserID, apiBody.ID)
		assert.Equal(t, body.UserName, apiBody.UserName)
		assert.Equal(t, body.GivenName, apiBody.Name.GivenName)
		assert.Equal(t, body.FamilyName, apiBody.Name.FamilyName)
		assert.Equal(t, body.Active, apiBody.Active)
	})

	t.Run("should return an error when passing an empty id to api replace user body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidReplaceUser()
		convertPorcelainToReplaceUserRequest("", body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the user id")
	})

	t.Run("should return an error when passing an empty userName to api replace user body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidReplaceUser()
		body.UserName = ""
		convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the user email")
	})

	t.Run("should return an error when passing an empty givenName to api replace user body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidReplaceUser()
		body.GivenName = ""
		convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the user first name")
	})

	t.Run("should return an error when passing an empty familyName to api replace user body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidReplaceUser()
		body.FamilyName = ""
		convertPorcelainToReplaceUserRequest(mockUserID, body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the user last name")
	})

	t.Run("should convert a create user body to api body when passing a valid create user body", func(t *testing.T) {
		body := getValidCreateUser()
		apiBody := convertPorcelainToCreateUserRequest(body)

		assert.Equal(t, body.UserName, apiBody.UserName)
		assert.Equal(t, body.GivenName, apiBody.Name.GivenName)
		assert.Equal(t, body.FamilyName, apiBody.Name.FamilyName)
		assert.Equal(t, body.Active, apiBody.Active)
	})

	t.Run("should return an error when passing an empty userName to api create user body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidCreateUser()
		body.UserName = ""
		convertPorcelainToCreateUserRequest(body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the user email")
	})

	t.Run("should return an error when passing an empty givenName to api create user body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidCreateUser()
		body.GivenName = ""
		convertPorcelainToCreateUserRequest(body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the user first name")
	})

	t.Run("should return an error when passing an empty familyName to api create user body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidCreateUser()
		body.FamilyName = ""
		convertPorcelainToCreateUserRequest(body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the user last name")
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
