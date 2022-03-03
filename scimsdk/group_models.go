package scimsdk

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
	Members     []GroupMember
}

type ReplaceGroupBody CreateGroupBody

type UpdateGroupReplaceName struct {
	DisplayName string
}
