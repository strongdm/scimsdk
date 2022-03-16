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

type GroupSmokeTest struct{}

func TestGroupSmoke(t *testing.T) {
	ExecuteSmokeTests(t, reflect.TypeOf(GroupSmokeTest{}), initializeSentry, flushSentry)
}

func (groupTest GroupSmokeTest) CommonFlow(t *testing.T) {
	var errors []AssertErr = []AssertErr{}
	defer sendErrorsToSentry(convertAssertErrListToStrList(errors))
	monkey.Patch(assert.Fail, func(t assert.TestingT, message string, msgAndArgs ...interface{}) bool {
		caller := getCaller()
		errors = append(errors, AssertErr{
			Message:      message,
			Caller:       caller,
			EntityParent: "Group",
		})
		return mockAssertFail(t, message, msgAndArgs...)
	})

	assert := assert.New(t)
	token := os.Getenv("SDM_SCIM_TOKEN")

	assert.NotEmpty(token)

	client := scimsdk.NewClient(token, nil)

	// Assert Create User Method
	user, err := client.Users().Create(context.Background(), scimsdk.CreateUser{
		UserName:   "user@email.com",
		GivenName:  "test",
		FamilyName: "name",
		Active:     true,
	})

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

	assert.Nil(err)
	assert.NotNil(group)
	assert.NotEmpty(group.ID)
	assert.NotEmpty(group.DisplayName)
	assert.NotNil(group.Meta)
	assert.NotEmpty(group.Meta.Location)
	assert.NotEmpty(group.Meta.ResourceType)
	assert.Equal(len(group.Members), 1)

	firstMember := group.Members[0]

	assert.NotEmpty(firstMember.ID)
	assert.NotEmpty(firstMember.Email)

	// Assert find Group method
	group, err = client.Groups().Find(context.Background(), group.ID)

	assert.Nil(err)
	assert.NotNil(group)
	assert.NotEmpty(group.ID)
	assert.NotEmpty(group.DisplayName)
	assert.NotNil(group.Meta)
	assert.NotEmpty(group.Meta.Location)
	assert.NotEmpty(group.Meta.ResourceType)
	assert.Equal(len(group.Members), 1)

	firstMember = group.Members[0]

	assert.NotEmpty(firstMember.ID)
	assert.NotEmpty(firstMember.Email)

	// TODO: ADD UPDATE GROUP METHOD ASSERTS

	// Assert Replace Group Method
	group, err = client.Groups().Replace(context.Background(), group.ID, scimsdk.ReplaceGroupBody{
		DisplayName: "Replaced Display Name",
		Members:     []scimsdk.GroupMember{},
	})

	assert.Nil(err)
	assert.NotNil(group)
	assert.NotEmpty(group.ID)
	assert.NotEmpty(group.DisplayName)
	assert.NotNil(group.Meta)
	assert.NotEmpty(group.Meta.Location)
	assert.NotEmpty(group.Meta.ResourceType)
	assert.Equal(len(group.Members), 0)

	ok, err := client.Groups().UpdateAddMembers(context.Background(), group.ID, []scimsdk.GroupMember{
		{
			ID:    user.ID,
			Email: user.UserName,
		},
	})

	assert.Nil(err)
	assert.True(ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assert.Nil(err)
	assert.NotNil(group)
	assert.NotEmpty(group.ID)
	assert.NotEmpty(group.DisplayName)
	assert.NotNil(group.Meta)
	assert.NotEmpty(group.Meta.Location)
	assert.NotEmpty(group.Meta.ResourceType)
	assert.Equal(len(group.Members), 1)

	firstMember = group.Members[0]

	assert.NotEmpty(firstMember.ID)
	assert.NotEmpty(firstMember.Email)

	ok, err = client.Groups().UpdateRemoveMemberByID(context.Background(), group.ID, user.ID)

	assert.Nil(err)
	assert.True(ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assert.Nil(err)
	assert.NotNil(group)
	assert.NotEmpty(group.ID)
	assert.NotEmpty(group.DisplayName)
	assert.NotNil(group.Meta)
	assert.NotEmpty(group.Meta.Location)
	assert.NotEmpty(group.Meta.ResourceType)
	assert.Equal(len(group.Members), 0)

	newGroupName := "New Group Name"
	ok, err = client.Groups().UpdateReplaceName(context.Background(), group.ID, scimsdk.UpdateGroupReplaceName{DisplayName: newGroupName})

	assert.Nil(err)
	assert.True(ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assert.Nil(err)
	assert.NotNil(group)
	assert.NotEmpty(group.ID)
	assert.Equal(group.DisplayName, newGroupName)
	assert.NotNil(group.Meta)
	assert.NotEmpty(group.Meta.Location)
	assert.NotEmpty(group.Meta.ResourceType)
	assert.Equal(len(group.Members), 0)

	secondUser, err := client.Users().Create(context.Background(), scimsdk.CreateUser{
		UserName:   "secondUser@email.com",
		GivenName:  "second test",
		FamilyName: "second name",
		Active:     true,
	})

	assert.Nil(err)
	assert.NotNil(secondUser)
	assert.NotEmpty(user.DisplayName)
	assert.Greater(len(user.Emails), 0)
	assert.Equal(len(user.Groups), 0)
	assert.NotEmpty(user.UserType)
	assert.NotEmpty(user.Name.FamilyName)
	assert.NotEmpty(user.Name.Formatted)
	assert.NotEmpty(user.Name.GivenName)
	assert.True(user.Active)

	ok, err = client.Groups().UpdateReplaceMembers(context.Background(), group.ID, []scimsdk.GroupMember{
		{
			ID:    secondUser.ID,
			Email: secondUser.UserName,
		},
	})

	assert.Nil(err)
	assert.True(ok)

	group, err = client.Groups().Find(context.Background(), group.ID)

	assert.Nil(err)
	assert.NotNil(group)
	assert.NotEmpty(group.ID)
	assert.Equal(group.DisplayName, newGroupName)
	assert.NotNil(group.Meta)
	assert.NotEmpty(group.Meta.Location)
	assert.NotEmpty(group.Meta.ResourceType)
	assert.Equal(len(group.Members), 1)

	firstMember = group.Members[0]

	assert.Equal(firstMember.ID, secondUser.ID)
	assert.Equal(firstMember.Email, secondUser.UserName)

	// Assert Delete Group Method
	ok, err = client.Groups().Delete(context.Background(), group.ID)

	assert.Nil(err)
	assert.True(ok)

	// Assert List Groups Method
	iterator := client.Groups().List(context.Background(), nil)

	assert.Empty(iterator.Err())

	for iterator.Next() {
		err := iterator.Err()
		group = iterator.Value()

		assert.Empty(err)
		assert.NotNil(group)
		assert.NotEmpty(group.ID)
		assert.NotEmpty(group.DisplayName)
		assert.NotNil(group.Meta)
		assert.NotEmpty(group.Meta.Location)
		assert.NotEmpty(group.Meta.ResourceType)

		if len(group.Members) > 0 {
			for _, member := range group.Members {
				assert.NotEmpty(member.ID)
				assert.NotEmpty(member.Email)
			}
		}
	}

	ok, err = client.Users().Delete(context.Background(), user.ID)

	assert.Nil(err)
	assert.True(ok)

	ok, err = client.Users().Delete(context.Background(), secondUser.ID)

	assert.Nil(err)
	assert.True(ok)
}
