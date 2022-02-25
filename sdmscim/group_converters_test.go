package sdmscim

import (
	"log"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestConvertGroupToAndFromPorcelain(t *testing.T) {
	t.Run("should convert a replace group body to api body when passing a valid replace group body", func(t *testing.T) {
		body := getValidReplaceGroup()
		apiBody := convertPorcelainToReplaceGroupRequest(body)

		firstMember := body.Members[0]
		firstApiMember := apiBody.Members[0]
		assert.Equal(t, body.DisplayName, apiBody.DisplayName)
		assert.Equal(t, len(body.Members), len(apiBody.Members))
		assert.Equal(t, firstApiMember.Display, firstMember.Display)
		assert.Equal(t, firstApiMember.Value, firstMember.Value)
	})

	t.Run("should return an error when passing an empty id to api replace group body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidReplaceGroup()
		body.DisplayName = ""
		convertPorcelainToReplaceGroupRequest(body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the group display name")
	})

	t.Run("should convert a create group body to api body when passing a valid create group body", func(t *testing.T) {
		body := getValidCreateGroup()
		apiBody := convertPorcelainToCreateGroupRequest(body)

		firstMember := body.Members[0]
		firstApiMember := apiBody.Members[0]
		assert.Equal(t, body.DisplayName, apiBody.DisplayName)
		assert.Equal(t, len(body.Members), len(apiBody.Members))
		assert.Equal(t, firstApiMember.Display, firstMember.Display)
		assert.Equal(t, firstApiMember.Value, firstMember.Value)
	})

	t.Run("should return an error when passing an empty displayName to api create group body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidCreateGroup()
		body.DisplayName = ""
		convertPorcelainToCreateGroupRequest(body)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the group display name")
	})

	t.Run("should convert a group member list to api group members list when passing a valid group member list", func(t *testing.T) {
		groupDisplay := "yyy"
		groupValue := "xxx"
		members := []GroupMember{{Display: groupDisplay, Value: groupValue}}
		apiBody := convertPorcelainToCreateMembersRequest(members)

		firstApiMember := apiBody[0]
		assert.NotNil(t, apiBody)
		assert.Equal(t, firstApiMember.Display, groupDisplay)
		assert.Equal(t, firstApiMember.Value, groupValue)
	})

	t.Run("should return an error when passing an empty member display to api group member", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		groupValue := "xxx"
		members := []GroupMember{{Value: groupValue}}
		convertPorcelainToCreateMembersRequest(members)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the member display")
	})

	t.Run("should return an error when passing an empty member value to api group member", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		groupDisplay := "yyy"
		members := []GroupMember{{Display: groupDisplay}}
		convertPorcelainToCreateMembersRequest(members)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the member value")
	})

	t.Run("should convert a group member list to api group member list when passing a valid group member list", func(t *testing.T) {
		groupDisplay := "yyy"
		groupValue := "xxx"
		members := []GroupMember{{Display: groupDisplay, Value: groupValue}}
		apiBody := convertPorcelainToUpdateGroupReplaceMembers(members)

		firstApiMember := apiBody.Operations[0].(apiUpdateGroupOperationRequest).Value.([]apiGroupMemberRequest)[0]
		assert.NotNil(t, apiBody)
		assert.Equal(t, firstApiMember.Display, groupDisplay)
		assert.Equal(t, firstApiMember.Value, groupValue)
	})

	t.Run("should return an error when passing an empty member display to api replace group members", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidCreateGroup()
		body.Members[0].Display = ""
		convertPorcelainToUpdateGroupReplaceMembers(body.Members)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the member display")
	})

	t.Run("should return an error when passing an empty member value to api replace group members", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		body := getValidCreateGroup()
		body.Members[0].Value = ""
		convertPorcelainToUpdateGroupReplaceMembers(body.Members)

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the member value")
	})

	t.Run("should convert a group replace name body to api group replace name body when passing a valid group name", func(t *testing.T) {
		groupName := "group name"
		apiBody := convertPorcelainToUpdateGroupName(UpdateGroupReplaceName{DisplayName: groupName})

		assert.NotNil(t, apiBody)
		operation := apiBody.Operations[0].(apiUpdateGroupOperationRequest)
		mappedOperationValue := operation.Value.(map[string]string)
		assert.Equal(t, groupName, mappedOperationValue["displayName"])
	})

	t.Run("should return an error when passing an empty group name to api group replace name body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		convertPorcelainToUpdateGroupName(UpdateGroupReplaceName{})

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the group name")
	})

	t.Run("should convert a member id to api group remove member body when passing a valid member id", func(t *testing.T) {
		memberID := "user-xxx"
		apiBody := convertPorcelainToUpdateGroupRemoveMember(memberID)

		assert.NotNil(t, apiBody)
		operation := apiBody.Operations[0].(*apiUpdateGroupOperationRequest)
		assert.Contains(t, operation.Path, memberID)
	})

	t.Run("should return an error when passing an empty member id to api group remove member body", func(t *testing.T) {
		exitStatus := 0
		fatalMessage := ""
		monkey.Patch(log.Fatal, func(args ...interface{}) {
			exitStatus = 1
			fatalMessage = args[0].(string)
		})
		convertPorcelainToUpdateGroupRemoveMember("")

		assert.Equal(t, exitStatus, 1)
		assert.Contains(t, fatalMessage, "must pass the member id")
	})
}
