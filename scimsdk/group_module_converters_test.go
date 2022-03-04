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

		firstMember := body.Members[0]
		firstApiMember := apiBody.Members[0]
		assert.Nil(t, err)
		assert.Equal(t, body.DisplayName, apiBody.DisplayName)
		assert.Equal(t, len(body.Members), len(apiBody.Members))
		assert.Equal(t, firstApiMember.Display, firstMember.Display)
		assert.Equal(t, firstApiMember.Value, firstMember.Value)
	})

	t.Run("should return an error when passing an empty id to api replace group body", func(t *testing.T) {
		body := getValidReplaceGroup()
		body.DisplayName = ""
		_, err := convertPorcelainToReplaceGroupRequest(body)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "must pass the group display name")
	})

	t.Run("should convert a create group body to api body when passing a valid create group body", func(t *testing.T) {
		body := getValidCreateGroup()
		apiBody, err := convertPorcelainToCreateGroupRequest(body)

		firstMember := body.Members[0]
		firstApiMember := apiBody.Members[0]
		assert.Nil(t, err)
		assert.Equal(t, body.DisplayName, apiBody.DisplayName)
		assert.Equal(t, len(body.Members), len(apiBody.Members))
		assert.Equal(t, firstApiMember.Display, firstMember.Display)
		assert.Equal(t, firstApiMember.Value, firstMember.Value)
	})

	t.Run("should return an error when passing an empty displayName to api create group body", func(t *testing.T) {
		body := getValidCreateGroup()
		body.DisplayName = ""
		_, err := convertPorcelainToCreateGroupRequest(body)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "must pass the group display name")
	})

	t.Run("should convert a group member list to api group members list when passing a valid group member list", func(t *testing.T) {
		groupDisplay := "yyy"
		groupValue := "xxx"
		members := []GroupMember{{Display: groupDisplay, Value: groupValue}}
		apiBody, err := convertPorcelainToCreateMembersRequest(members)

		firstApiMember := apiBody[0]
		assert.Nil(t, err)
		assert.NotNil(t, apiBody)
		assert.Equal(t, firstApiMember.Display, groupDisplay)
		assert.Equal(t, firstApiMember.Value, groupValue)
	})

	t.Run("should return an error when passing an empty member display to api group member", func(t *testing.T) {
		groupValue := "xxx"
		members := []GroupMember{{Value: groupValue}}
		_, err := convertPorcelainToCreateMembersRequest(members)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the member display")
	})

	t.Run("should return an error when passing an empty member value to api group member", func(t *testing.T) {
		groupDisplay := "yyy"
		members := []GroupMember{{Display: groupDisplay}}
		_, err := convertPorcelainToCreateMembersRequest(members)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the member value")
	})

	t.Run("should convert a group member list to api group member list when passing a valid group member list", func(t *testing.T) {
		groupDisplay := "yyy"
		groupValue := "xxx"
		members := []GroupMember{{Display: groupDisplay, Value: groupValue}}
		apiBody, err := convertPorcelainToUpdateGroupReplaceMembersRequest(members)

		firstApiMember := apiBody.Operations[0].(service.UpdateGroupOperationRequest).Value.([]service.GroupMemberRequest)[0]
		assert.Nil(t, err)
		assert.NotNil(t, apiBody)
		assert.Equal(t, firstApiMember.Display, groupDisplay)
		assert.Equal(t, firstApiMember.Value, groupValue)
	})

	t.Run("should return an error when passing an empty member display to api replace group members", func(t *testing.T) {
		body := getValidCreateGroup()
		body.Members[0].Display = ""
		_, err := convertPorcelainToUpdateGroupReplaceMembersRequest(body.Members)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the member display")
	})

	t.Run("should return an error when passing an empty member value to api replace group members", func(t *testing.T) {
		body := getValidCreateGroup()
		body.Members[0].Value = ""
		_, err := convertPorcelainToUpdateGroupReplaceMembersRequest(body.Members)

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the member value")
	})

	t.Run("should convert a group replace name body to api group replace name body when passing a valid group name", func(t *testing.T) {
		groupName := "group name"
		apiBody, err := convertPorcelainToUpdateGroupNameRequest(UpdateGroupReplaceName{DisplayName: groupName})

		assert.Nil(t, err)
		assert.NotNil(t, apiBody)
		operation := apiBody.Operations[0].(service.UpdateGroupOperationRequest)
		mappedOperationValue := operation.Value.(map[string]string)
		assert.Equal(t, groupName, mappedOperationValue["displayName"])
	})

	t.Run("should return an error when passing an empty group name to api group replace name body", func(t *testing.T) {
		_, err := convertPorcelainToUpdateGroupNameRequest(UpdateGroupReplaceName{})

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the group name")
	})

	t.Run("should convert a member id to api group remove member body when passing a valid member id", func(t *testing.T) {
		memberID := "user-xxx"
		apiBody, err := convertPorcelainToUpdateGroupRemoveMemberRequest(memberID)

		assert.Nil(t, err)
		assert.NotNil(t, apiBody)
		operation := apiBody.Operations[0].(*service.UpdateGroupOperationRequest)
		assert.Contains(t, operation.Path, memberID)
	})

	t.Run("should return an error when passing an empty member id to api group remove member body", func(t *testing.T) {
		_, err := convertPorcelainToUpdateGroupRemoveMemberRequest("")

		assert.NotNil(t, err, 1)
		assert.Contains(t, err.Error(), "must pass the member id")
	})
}

func getValidCreateGroup() *CreateGroupBody {
	return &CreateGroupBody{
		DisplayName: "xxx",
		Members: []GroupMember{
			{
				Display: "xxx",
				Value:   "yyy",
			},
		},
	}
}

func getValidReplaceGroup() *ReplaceGroupBody {
	return &ReplaceGroupBody{
		DisplayName: "xxx",
		Members: []GroupMember{
			{
				Display: "zzz",
				Value:   "www",
			},
		},
	}
}
