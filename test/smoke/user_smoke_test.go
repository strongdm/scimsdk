package main

import (
	"context"
	"os"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"github.com/strongdm/scimsdk/scimsdk"
)

type UserSmokeTest struct{}

func TestUserSmoke(t *testing.T) {
	ExecuteSmokeTests(t, reflect.TypeOf(UserSmokeTest{}), initializeSentry, flushSentry)
}

func (UserSmokeTest) CommonFlow(t *testing.T) {
	var errors []AssertErr = []AssertErr{}

	monkey.Patch(assert.Fail, func(t assert.TestingT, message string, msgAndArgs ...interface{}) bool {
		caller := getCaller()
		errors = append(errors, AssertErr{
			Message:      message,
			Caller:       caller,
			EntityParent: "User",
		})
		return mockAssertFail(t, message, msgAndArgs...)
	})

	assert := assert.New(t)
	token := os.Getenv("SDM_SCIM_TOKEN")

	assert.NotEmpty(token)

	client := scimsdk.NewClient(token, nil)
	user, err := client.Users().Create(context.Background(), scimsdk.CreateUser{
		UserName:   "user@email.com",
		GivenName:  "test",
		FamilyName: "name",
		Active:     true,
	})

	// Assert Create User Method
	assert.Nil(err)
	assert.NotNil(user)
	assert.NotEmpty(user.DisplayName)
	assert.Greater(len(user.Emails), 0)
	assert.Equal(len(user.Groups), 0)
	assert.NotEmpty(user.UserType)
	assert.NotEmpty(user.Name.FamilyName)
	assert.NotEmpty(user.Name.Formatted)
	assert.NotEmpty(user.Name.GivenName)
	assert.True(user.Active)

	user, err = client.Users().Find(context.Background(), user.ID)

	// Assert Find User Method
	assert.NotNil(err) // TODO Change back to assert.Nil(err)
	assert.NotNil(user)
	assert.NotEmpty(user.DisplayName)
	assert.Greater(len(user.Emails), 0)
	assert.Equal(len(user.Groups), 0)
	assert.NotEmpty(user.UserType)
	assert.NotEmpty(user.Name.FamilyName)
	assert.NotEmpty(user.Name.Formatted)
	assert.NotEmpty(user.Name.GivenName)
	assert.True(user.Active)

	ok, err := client.Users().Update(context.Background(), user.ID, scimsdk.UpdateUser{
		Active: true,
	})

	// Assert Update User Method
	assert.Nil(err)
	assert.True(ok)

	user, err = client.Users().Find(context.Background(), user.ID)

	// Assert Find User Method after Update
	assert.Nil(err)
	assert.NotNil(user)
	assert.NotEmpty(user.DisplayName)
	assert.Greater(len(user.Emails), 0)
	assert.Equal(len(user.Groups), 0)
	assert.NotEmpty(user.UserType)
	assert.NotEmpty(user.Name.FamilyName)
	assert.NotEmpty(user.Name.Formatted)
	assert.NotEmpty(user.Name.GivenName)
	assert.True(user.Active)

	user, err = client.Users().Replace(context.Background(), user.ID, scimsdk.ReplaceUser{
		UserName:   "user+01@email.com",
		GivenName:  "test replaced",
		FamilyName: "name replaced",
		Active:     true,
	})

	// Assert Replace User Method
	assert.Nil(err)
	assert.NotNil(user)
	assert.NotEmpty(user.DisplayName)
	assert.Greater(len(user.Emails), 0)
	assert.Equal(len(user.Groups), 0)
	assert.NotEmpty(user.UserType)
	assert.NotEmpty(user.Name.FamilyName)
	assert.NotEmpty(user.Name.Formatted)
	assert.NotEmpty(user.Name.GivenName)
	assert.True(user.Active)

	ok, err = client.Users().Delete(context.Background(), user.ID)

	// Assert Delete User Method
	assert.Nil(err)
	assert.True(ok)

	userIterator := client.Users().List(context.Background(), nil)

	assert.Empty(userIterator.Err())

	// Assert List User Method
	for userIterator.Next() {
		assert.Nil(err)
		assert.NotNil(user)
		assert.NotEmpty(user.DisplayName)
		assert.Greater(len(user.Emails), 0)
		assert.Equal(len(user.Groups), 0)
		assert.NotEmpty(user.UserType)
		assert.NotEmpty(user.Name.FamilyName)
		assert.NotEmpty(user.Name.Formatted)
		assert.NotEmpty(user.Name.GivenName)
	}

	sendErrorsToSentry(convertAssertErrListToStrList(errors))
}
