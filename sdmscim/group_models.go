package sdmscim

type Group struct {
	ID          string
	DisplayName string
	Members     []*GroupMember
	Meta        *GroupMetadata
}

type GroupMember struct {
	Value   string
	Display string
}

type GroupMetadata struct {
	ResourceType string
	Location     string
}

type CreateGroupBody struct {
	DisplayName string
	Members     []*GroupMember
}

type ReplaceGroupBody CreateGroupBody

// Add members
/*
	{
		"op": "add",
		"path": "members",
		"value": [
			{"value":"a-0001",
			"display":"myUser@example.test"}
		]
	}
*/
type UpdateGroupMemberBody struct {
	Value   string
	Display string
}

// Replace group name
/*
	{
		"op": "replace",
		"value": {
			"displayName: "newName"
		}
	}
*/
type UpdateGroupReplaceNameBody struct {
	Operations []UpdateGroupReplaceNameOperationBody
}
type UpdateGroupReplaceNameOperationBody struct {
	OP    string
	Value []UpdateGroupReplaceNameOperationValueBody
}
type UpdateGroupReplaceNameOperationValueBody struct {
	DisplayName string
}

// Replace group members
/*
	{
		"op": "replace",
		"path": "members",
		"value": [
			{
				"value":"a-0001",
				"display":"myUser@example.test"
			}
		]
	}
*/
type UpdateGroupReplaceMembersBody struct {
	Operations []UpdateGroupReplaceMembersOperationBody
}
type UpdateGroupReplaceMembersOperationBody struct {
	OP    string
	Path  string
	Value []UpdateGroupReplaceMembersOperationValueBody
}
type UpdateGroupReplaceMembersOperationValueBody struct {
	Value   string
	Display string
}
