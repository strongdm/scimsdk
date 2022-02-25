package sdmscim

type apiGroupPageResponse struct {
	Resources    []*apiGroupResponse `json:"Resources"`
	ItemsPerPage int                 `json:"itemsPerPage"`
	Schemas      []string            `json:"schemas"`
	StartIndex   int                 `json:"startIndex"`
	TotalResults int                 `json:"totalResults"`
}

type apiGroupResponse struct {
	Schemas     []string                  `json:"schemas"`
	DisplayName string                    `json:"displayName"`
	ID          string                    `json:"id"`
	Members     []*apiGroupMemberResponse `json:"members"`
	Meta        *apiGroupMetadataResponse `json:"meta"`
}

type apiGroupMemberResponse struct {
	Value   string `json:"value"`
	Display string `json:"display"`
}

type apiGroupMetadataResponse struct {
	ResourceType string `json:"resourceType"`
	Location     string `json:"location"`
}

type apiCreateGroupRequest struct {
	DisplayName string                    `json:"displayName"`
	Members     []*apiCreateMemberRequest `json:"members"`
	Schemas     []string                  `json:"schemas"`
}

type apiReplaceGroupRequest apiCreateGroupRequest

type apiCreateMemberRequest struct {
	Value   string `json:"value"`
	Display string `json:"display"`
}

type apiUpdateGroupRequest struct {
	Schemas    []string
	Operations []interface{} `json:"Operations"`
}

type apiUpdateGroupAddMembersOperationRequest struct {
	OP    string                                          `json:"op"`
	Path  string                                          `json:"path"`
	Value []apiUpdateGroupAddMembersOperationValueRequest `json:"value"`
}

type apiUpdateGroupAddMembersOperationValueRequest struct {
	Value   string `json:"value"`
	Display string `json:"display"`
}

type apiUpdateGroupReplaceNameOperationRequest struct {
	OP    string                                     `json:"op"`
	Value []UpdateGroupReplaceNameOperationValueBody `json:"value"`
}

type apiUpdateGroupReplaceNameOperationValueRequest struct {
	DisplayName string `json:"displayName"`
}

type apiUpdateGroupReplaceMembersOperationRequest struct {
	OP    string                                              `json:"op"`
	Path  string                                              `json:"path"`
	Value []apiUpdateGroupReplaceMembersOperationValueRequest `json:"value"`
}

type apiUpdateGroupReplaceMembersOperationValueRequest struct {
	Value   string `json:"value"`
	Display string `json:"display"`
}

type apiUpdateGroupRemoveMembersOperationRequest struct {
	OP   string `json:"op"`
	Path string `json:"path"`
}
