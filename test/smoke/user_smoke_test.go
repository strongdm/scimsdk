package main

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/strongdm/scimsdk"
	"github.com/strongdm/scimsdk/models"
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
	user, err := client.Users().Create(context.Background(), models.CreateUser{
		UserName:   os.Getenv("SDM_SCIM_TEST_USERNAME1"),
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
	// assertNotNil(t, err) 
	assertNotNil(t, user)
	assertNotEmpty(t, user.DisplayName)
	assertGreater(t, len(user.Emails), 0)
	assertEqual(t, len(user.Groups), 0)
	assertNotEmpty(t, user.UserType)
	assertNotEmpty(t, user.Name.FamilyName)
	assertNotEmpty(t, user.Name.Formatted)
	assertNotEmpty(t, user.Name.GivenName)
	assertTrue(t, user.Active)

	ok, err := client.Users().Update(context.Background(), user.ID, models.UpdateUser{
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

	user, err = client.Users().Replace(context.Background(), user.ID, models.ReplaceUser{
		UserName:   os.Getenv("SDM_SCIM_TEST_USERNAME2"),
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
