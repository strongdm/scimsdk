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
}
