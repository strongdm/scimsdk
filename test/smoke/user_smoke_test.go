package main

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/strongdm/scimsdk/scimsdk"
)

type UserSmokeTest struct{}

func TestUserSmoke(t *testing.T) {
	ExecuteSmokeTests(t, reflect.TypeOf(UserSmokeTest{}), initializeSentry, flushSentry)
}

func (UserSmokeTest) CommonFlow(t *testing.T) {
	defer sendErrorsToSentry()

	token := os.Getenv("SDM_SCIM_TOKEN")

	assertNotEmpty(t, token)

	client := scimsdk.NewClient(token, nil)

	// Assert Create User Method
	user, err := client.Users().Create(context.Background(), scimsdk.CreateUser{
		UserName:   "user@email.com",
		GivenName:  "test",
		FamilyName: "name",
		Active:     true,
	})

	// Assert Create User Method
	assertNil(t, err)
	assertNotNil(t, user)
	assertNotEmpty(t, user.DisplayName)
	assertGreater(t, len(user.Emails), 0)
	assertEqual(t, len(user.Groups), 0)
	assertNotEmpty(t, user.UserType)
	assertNotEmpty(t, user.Name.FamilyName)
	assertNotEmpty(t, user.Name.Formatted)
	assertNotEmpty(t, user.Name.GivenName)
	assertTrue(t, user.Active)

	user, err = client.Users().Find(context.Background(), user.ID)

	// Assert Find User Method
	assertNil(t, err)
	assertNotNil(t, user)
	assertNotEmpty(t, user.DisplayName)
	assertGreater(t, len(user.Emails), 0)
	assertEqual(t, len(user.Groups), 0)
	assertNotEmpty(t, user.UserType)
	assertNotEmpty(t, user.Name.FamilyName)
	assertNotEmpty(t, user.Name.Formatted)
	assertNotEmpty(t, user.Name.GivenName)
	assertTrue(t, user.Active)

	ok, err := client.Users().Update(context.Background(), user.ID, scimsdk.UpdateUser{
		Active: true,
	})

	// Assert Update User Method
	assertNil(t, err)
	assertTrue(t, ok)

	user, err = client.Users().Find(context.Background(), user.ID)

	// Assert Find User Method after Update
	assertNil(t, err)
	assertNotNil(t, user)
	assertNotEmpty(t, user.DisplayName)
	assertGreater(t, len(user.Emails), 0)
	assertEqual(t, len(user.Groups), 0)
	assertNotEmpty(t, user.UserType)
	assertNotEmpty(t, user.Name.FamilyName)
	assertNotEmpty(t, user.Name.Formatted)
	assertNotEmpty(t, user.Name.GivenName)
	assertTrue(t, user.Active)

	user, err = client.Users().Replace(context.Background(), user.ID, scimsdk.ReplaceUser{
		UserName:   "user+01@email.com",
		GivenName:  "test replaced",
		FamilyName: "name replaced",
		Active:     true,
	})

	// Assert Replace User Method
	assertNil(t, err)
	assertNotNil(t, user)
	assertNotEmpty(t, user.DisplayName)
	assertGreater(t, len(user.Emails), 0)
	assertEqual(t, len(user.Groups), 0)
	assertNotEmpty(t, user.UserType)
	assertNotEmpty(t, user.Name.FamilyName)
	assertNotEmpty(t, user.Name.Formatted)
	assertNotEmpty(t, user.Name.GivenName)
	assertTrue(t, user.Active)

	ok, err = client.Users().Delete(context.Background(), user.ID)

	// Assert Delete User Method
	assertNil(t, err)
	assertTrue(t, ok)

	userIterator := client.Users().List(context.Background(), nil)

	assertEmpty(t, userIterator.Err())

	// Assert List User Method
	for userIterator.Next() {
		assertNil(t, err)
		assertNotNil(t, user)
		assertNotEmpty(t, user.DisplayName)
		assertGreater(t, len(user.Emails), 0)
		assertEqual(t, len(user.Groups), 0)
		assertNotEmpty(t, user.UserType)
		assertNotEmpty(t, user.Name.FamilyName)
		assertNotEmpty(t, user.Name.Formatted)
		assertNotEmpty(t, user.Name.GivenName)
	}
}
