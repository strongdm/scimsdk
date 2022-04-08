package main

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/strongdm/scimsdk/scimsdk"
)

type GroupSmokeTest struct{}

func TestGroupSmoke(t *testing.T) {
	ExecuteSmokeTests(t, reflect.TypeOf(GroupSmokeTest{}), initializeSentry, flushSentry)
}

func (groupTest GroupSmokeTest) CommonFlow(t *testing.T) {
	defer sendErrorsToSentry()

	token := os.Getenv("SDM_SCIM_TOKEN")

	assertNotEmpty(t, token)

	client := scimsdk.NewClient(token, nil)

	// Assert Create User Method
	user, err := client.Users().Create(context.Background(), scimsdk.CreateUser{
		UserName:   os.Getenv("SDM_SCIM_TEST_USERNAME1"),
		GivenName:  "test",
		FamilyName: "name",
		Active:     true,
	})

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

	// Assert Create Group Method
	group, err := client.Groups().Create(context.Background(), scimsdk.CreateGroupBody{
		DisplayName: "xxx",
		Members: []scimsdk.GroupMember{
			{
				ID:    user.ID,
				Email: user.UserName,
			},
		},
	})

	assertNil(t, err)
	assertNotNil(t, group)
	assertNotEmpty(t, group.ID)
	assertNotEmpty(t, group.DisplayName)
	assertNotNil(t, group.Meta)
	assertNotEmpty(t, group.Meta.Location)
	assertNotEmpty(t, group.Meta.ResourceType)
	assertEqual(t, len(group.Members), 1)

	firstMember := group.Members[0]

	assertNotEmpty(t, firstMember.ID)
	assertNotEmpty(t, firstMember.Email)

	// Assert find Group method
	group, err = client.Groups().Find(context.Background(), group.ID)

	assertNil(t, err)
	assertNotNil(t, group)
	assertNotEmpty(t, group.ID)
	assertNotEmpty(t, group.DisplayName)
	assertNotNil(t, group.Meta)
	assertNotEmpty(t, group.Meta.Location)
	assertNotEmpty(t, group.Meta.ResourceType)
	assertEqual(t, len(group.Members), 1)

	firstMember = group.Members[0]

	assertNotEmpty(t, firstMember.ID)
	assertNotEmpty(t, firstMember.Email)

	// Assert Replace Group Method
	group, err = client.Groups().Replace(context.Background(), group.ID, scimsdk.ReplaceGroupBody{
		DisplayName: "Replaced Display Name",
		Members:     []scimsdk.GroupMember{},
	})

	assertNil(t, err)
	assertNotNil(t, group)
	assertNotEmpty(t, group.ID)
	assertNotEmpty(t, group.DisplayName)
	assertNotNil(t, group.Meta)
	assertNotEmpty(t, group.Meta.Location)
	assertNotEmpty(t, group.Meta.ResourceType)
	assertEqual(t, len(group.Members), 0)

	ok, err := client.Groups().UpdateAddMembers(context.Background(), group.ID, []scimsdk.GroupMember{
		{
			ID:    user.ID,
			Email: user.UserName,
		},
	})

	assertNil(t, err)
	assertTrue(t, ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assertNil(t, err)
	assertNotNil(t, group)
	assertNotEmpty(t, group.ID)
	assertNotEmpty(t, group.DisplayName)
	assertNotNil(t, group.Meta)
	assertNotEmpty(t, group.Meta.Location)
	assertNotEmpty(t, group.Meta.ResourceType)
	assertEqual(t, len(group.Members), 1)

	firstMember = group.Members[0]

	assertNotEmpty(t, firstMember.ID)
	assertNotEmpty(t, firstMember.Email)

	ok, err = client.Groups().UpdateRemoveMemberByID(context.Background(), group.ID, user.ID)

	assertNil(t, err)
	assertTrue(t, ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assertNil(t, err)
	assertNotNil(t, group)
	assertNotEmpty(t, group.ID)
	assertNotEmpty(t, group.DisplayName)
	assertNotNil(t, group.Meta)
	assertNotEmpty(t, group.Meta.Location)
	assertNotEmpty(t, group.Meta.ResourceType)
	assertEqual(t, len(group.Members), 0)

	newGroupName := "New Group Name"
	ok, err = client.Groups().UpdateReplaceName(context.Background(), group.ID, scimsdk.UpdateGroupReplaceName{DisplayName: newGroupName})

	assertNil(t, err)
	assertTrue(t, ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assertNil(t, err)
	assertNotNil(t, group)
	assertNotEmpty(t, group.ID)
	assertEqual(t, group.DisplayName, newGroupName)
	assertNotNil(t, group.Meta)
	assertNotEmpty(t, group.Meta.Location)
	assertNotEmpty(t, group.Meta.ResourceType)
	assertEqual(t, len(group.Members), 0)

	secondUser, err := client.Users().Create(context.Background(), scimsdk.CreateUser{
		UserName:   os.Getenv("SDM_SCIM_TEST_USERNAME2"),
		GivenName:  "second test",
		FamilyName: "second name",
		Active:     true,
	})

	assertNil(t, err)
	assertNotNil(t, secondUser)
	assertNotEmpty(t, user.DisplayName)
	assertGreater(t, len(user.Emails), 0)
	assertEqual(t, len(user.Groups), 0)
	assertNotEmpty(t, user.UserType)
	assertNotEmpty(t, user.Name.FamilyName)
	assertNotEmpty(t, user.Name.Formatted)
	assertNotEmpty(t, user.Name.GivenName)
	assertTrue(t, user.Active)

	ok, err = client.Groups().UpdateReplaceMembers(context.Background(), group.ID, []scimsdk.GroupMember{
		{
			ID:    secondUser.ID,
			Email: secondUser.UserName,
		},
	})

	assertNil(t, err)
	assertTrue(t, ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assertNil(t, err)
	assertNotNil(t, group)
	assertNotEmpty(t, group.ID)
	assertEqual(t, group.DisplayName, newGroupName)
	assertNotNil(t, group.Meta)
	assertNotEmpty(t, group.Meta.Location)
	assertNotEmpty(t, group.Meta.ResourceType)
	assertEqual(t, len(group.Members), 1)

	firstMember = group.Members[0]

	assertEqual(t, firstMember.ID, secondUser.ID)
	assertEqual(t, firstMember.Email, secondUser.UserName)

	// Assert Delete Group Method
	ok, err = client.Groups().Delete(context.Background(), group.ID)

	assertNil(t, err)
	assertTrue(t, ok)

	// Assert List Groups Method
	iterator := client.Groups().List(context.Background(), nil)

	assertEmpty(t, iterator.Err())

	for iterator.Next() {
		err := iterator.Err()
		group = iterator.Value()

		assertEmpty(t, err)
		assertNotNil(t, group)
		assertNotEmpty(t, group.ID)
		assertNotEmpty(t, group.DisplayName)
		assertNotNil(t, group.Meta)
		assertNotEmpty(t, group.Meta.Location)
		assertNotEmpty(t, group.Meta.ResourceType)

		if len(group.Members) > 0 {
			for _, member := range group.Members {
				assertNotEmpty(t, member.ID)
				assertNotEmpty(t, member.Email)
			}
		}
	}

	ok, err = client.Users().Delete(context.Background(), user.ID)

	assertNil(t, err)
	assertTrue(t, ok)

	ok, err = client.Users().Delete(context.Background(), secondUser.ID)

	assertNil(t, err)
	assertTrue(t, ok)
}
