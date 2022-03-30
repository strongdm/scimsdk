package scimsdk

import (
	"testing"

	"github.com/strongdm/scimsdk/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestConvertGroupToAndFromPorcelain(t *testing.T) {
	t.Run("should convert a replace group body to api body when passing a valid replace group body", func(t *testing.T) {
		body := getValidReplaceGroup()
		apiBody, err := convertPorcelainToReplaceGroupRequest(body)
		assertT := assert.New(t)

		firstMember := body.Members[0]
		firstApiMember := apiBody.Members[0]
		assertT.Nil(err)
		assertT.Equal(body.DisplayName, apiBody.DisplayName)
		assertT.Equal(len(body.Members), len(apiBody.Members))
		assertT.Equal(firstApiMember.Display, firstMember.Email)
		assertT.Equal(firstApiMember.Value, firstMember.ID)
	})

	t.Run("should return an error when passing an empty id to api replace group body", func(t *testing.T) {
		body := getValidReplaceGroup()
		body.DisplayName = ""
		_, err := convertPorcelainToReplaceGroupRequest(body)
		assertT := assert.New(t)

		assertT.NotNil(err)
		assertT.Contains(err.Error(), "must pass the group display name")
	})

	t.Run("should convert a create group body to api body when passing a valid create group body", func(t *testing.T) {
		body := getValidCreateGroup()
		apiBody, err := convertPorcelainToCreateGroupRequest(body)
		assertT := assert.New(t)

		firstMember := body.Members[0]
		firstApiMember := apiBody.Members[0]
		assertT.Nil(err)
		assertT.Equal(body.DisplayName, apiBody.DisplayName)
		assertT.Equal(len(body.Members), len(apiBody.Members))
		assertT.Equal(firstApiMember.Display, firstMember.Email)
		assertT.Equal(firstApiMember.Value, firstMember.ID)
	})

	t.Run("should return an error when passing an empty displayName to api create group body", func(t *testing.T) {
		body := getValidCreateGroup()
		body.DisplayName = ""
		_, err := convertPorcelainToCreateGroupRequest(body)
		assertT := assert.New(t)

		assertT.NotNil(err)
		assertT.Contains(err.Error(), "must pass the group display name")
	})

	t.Run("should convert a group member list to api group members list when passing a valid group member list", func(t *testing.T) {
		groupDisplay := "yyy"
		groupValue := "xxx"
		members := []GroupMember{{Email: groupDisplay, ID: groupValue}}
		apiBody, err := convertPorcelainToCreateMembersRequest(members)
		assertT := assert.New(t)

		firstApiMember := apiBody[0]
		assertT.Nil(err)
		assertT.NotNil(apiBody)
		assertT.Equal(firstApiMember.Display, groupDisplay)
		assertT.Equal(firstApiMember.Value, groupValue)
	})

	t.Run("should return an error when passing an empty member display to api group member", func(t *testing.T) {
		groupValue := "xxx"
		members := []GroupMember{{ID: groupValue}}
		_, err := convertPorcelainToCreateMembersRequest(members)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the member display")
	})

	t.Run("should return an error when passing an empty member value to api group member", func(t *testing.T) {
		groupDisplay := "yyy"
		members := []GroupMember{{Email: groupDisplay}}
		_, err := convertPorcelainToCreateMembersRequest(members)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the member value")
	})

	t.Run("should convert a group member list to api group member list when passing a valid group member list", func(t *testing.T) {
		groupDisplay := "yyy"
		groupValue := "xxx"
		members := []GroupMember{{Email: groupDisplay, ID: groupValue}}
		apiBody, err := convertPorcelainToUpdateGroupReplaceMembersRequest(members)
		assertT := assert.New(t)

		firstApiMember := apiBody.Operations[0].(service.UpdateGroupOperationRequest).Value.([]service.GroupMemberRequest)[0]
		assertT.Nil(err)
		assertT.NotNil(apiBody)
		assertT.Equal(firstApiMember.Display, groupDisplay)
		assertT.Equal(firstApiMember.Value, groupValue)
	})

	t.Run("should return an error when passing an empty member display to api replace group members", func(t *testing.T) {
		body := getValidCreateGroup()
		body.Members[0].Email = ""
		_, err := convertPorcelainToUpdateGroupReplaceMembersRequest(body.Members)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the member display")
	})

	t.Run("should return an error when passing an empty member value to api replace group members", func(t *testing.T) {
		body := getValidCreateGroup()
		body.Members[0].ID = ""
		_, err := convertPorcelainToUpdateGroupReplaceMembersRequest(body.Members)
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the member value")
	})

	t.Run("should convert a group replace name body to api group replace name body when passing a valid group name", func(t *testing.T) {
		groupName := "group name"
		apiBody, err := convertPorcelainToUpdateGroupNameRequest(UpdateGroupReplaceName{DisplayName: groupName})
		assertT := assert.New(t)

		assertT.Nil(err)
		assertT.NotNil(apiBody)
		operation := apiBody.Operations[0].(service.UpdateGroupOperationRequest)
		mappedOperationValue := operation.Value.(map[string]string)
		assertT.Equal(groupName, mappedOperationValue["displayName"])
	})

	t.Run("should return an error when passing an empty group name to api group replace name body", func(t *testing.T) {
		_, err := convertPorcelainToUpdateGroupNameRequest(UpdateGroupReplaceName{})
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the group name")
	})

	t.Run("should convert a member id to api group remove member body when passing a valid member id", func(t *testing.T) {
		memberID := "user-xxx"
		apiBody, err := convertPorcelainToUpdateGroupRemoveMemberRequest(memberID)
		assertT := assert.New(t)

		assertT.Nil(err)
		assertT.NotNil(apiBody)
		operation := apiBody.Operations[0].(*service.UpdateGroupOperationRequest)
		assertT.Contains(operation.Path, memberID)
	})

	t.Run("should return an error when passing an empty member id to api group remove member body", func(t *testing.T) {
		_, err := convertPorcelainToUpdateGroupRemoveMemberRequest("")
		assertT := assert.New(t)

		assertT.NotNil(err, 1)
		assertT.Contains(err.Error(), "must pass the member id")
	})
}

func getValidCreateGroup() *CreateGroupBody {
	return &CreateGroupBody{
		DisplayName: "xxx",
		Members: []GroupMember{
			{
				Email: "xxx",
				ID:    "yyy",
			},
		},
	}
}

func getValidReplaceGroup() *ReplaceGroupBody {
	return &ReplaceGroupBody{
		DisplayName: "xxx",
		Members: []GroupMember{
			{
				Email: "zzz",
				ID:    "www",
			},
		},
	}
}
