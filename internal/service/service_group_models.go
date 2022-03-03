package service

type GroupPageResponse struct {
	Resources    []*GroupResponse `json:"Resources"`
	ItemsPerPage int              `json:"itemsPerPage"`
	Schemas      []string         `json:"schemas"`
	StartIndex   int              `json:"startIndex"`
	TotalResults int              `json:"totalResults"`
}

type GroupResponse struct {
	Schemas     []string               `json:"schemas"`
	DisplayName string                 `json:"displayName"`
	ID          string                 `json:"id"`
	Members     []*GroupMemberResponse `json:"members"`
	Meta        *GroupMetadataResponse `json:"meta"`
}

type GroupMemberResponse struct {
	Value   string `json:"value"`
	Display string `json:"display"`
}

type GroupMetadataResponse struct {
	ResourceType string `json:"resourceType"`
	Location     string `json:"location"`
}

type CreateGroupRequest struct {
	DisplayName string                `json:"displayName"`
	Members     []*GroupMemberRequest `json:"members"`
	Schemas     []string              `json:"schemas"`
}

type ReplaceGroupRequest CreateGroupRequest

type GroupMemberRequest struct {
	Value   string `json:"value"`
	Display string `json:"display"`
}

type UpdateGroupRequest struct {
	Schemas    []string
	Operations []interface{} `json:"Operations"`
}

type UpdateGroupOperationRequest struct {
	OP    string      `json:"op"`
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}
