package sdmscim

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

const defaultGroupSchema = "urn:ietf:params:scim:schemas:core:2.0:Group"
const defaultPatchSchema = "urn:ietf:params:scim:api:messages:2.0:PatchOp"

func unmarshalGroupPageResponse(body io.ReadCloser) (*apiGroupPageResponse, error) {
	unmarshedResponse := &apiGroupPageResponse{}
	buff, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buff, &unmarshedResponse)
	if err != nil {
		return nil, err
	}
	return unmarshedResponse, nil
}

func unmarshalGroupResponse(body io.ReadCloser) (*apiGroupResponse, error) {
	unmarshedResponse := &apiGroupResponse{}
	buff, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buff, &unmarshedResponse)
	if err != nil {
		return nil, err
	}
	return unmarshedResponse, nil
}

func convertGroupResponseListToPorcelain(groupListResponse []*apiGroupResponse) []*Group {
	groupList := make([]*Group, 0)
	for _, groupResponse := range groupListResponse {
		groupList = append(groupList, convertGroupResponseToPorcelain(groupResponse))
	}
	return groupList
}

func convertGroupResponseToPorcelain(groupResponse *apiGroupResponse) *Group {
	return &Group{
		ID:          groupResponse.ID,
		DisplayName: groupResponse.DisplayName,
		Members:     convertGroupMemberResponseListToPorcelain(groupResponse.Members),
		Meta:        convertGroupMetaResponseToPorcelain(groupResponse.Meta),
	}
}

func convertGroupMemberResponseListToPorcelain(memberListResponse []*apiGroupMemberResponse) []*GroupMember {
	memberList := make([]*GroupMember, 0)
	for _, memberResponse := range memberListResponse {
		memberList = append(memberList, convertGroupMemberResponseToPorcelain(memberResponse))
	}
	return memberList
}

func convertGroupMemberResponseToPorcelain(memberResponse *apiGroupMemberResponse) *GroupMember {
	return &GroupMember{
		Value:   memberResponse.Value,
		Display: memberResponse.Display,
	}
}

func convertGroupMetaResponseToPorcelain(metaResponse *apiGroupMetadataResponse) *GroupMetadata {
	return &GroupMetadata{
		ResourceType: metaResponse.ResourceType,
		Location:     metaResponse.Location,
	}
}

func convertPorcelainToCreateGroupRequest(group *CreateGroupBody) *apiCreateGroupRequest {
	if group.DisplayName == "" {
		log.Fatal("You must pass the group display name in DisplayName field.")
	}
	return &apiCreateGroupRequest{
		Schemas:     []string{defaultGroupSchema},
		DisplayName: group.DisplayName,
		Members:     convertPorcelainToCreateMembersRequest(group.Members),
	}
}

func convertPorcelainToReplaceGroupRequest(group *ReplaceGroupBody) *apiReplaceGroupRequest {
	if group.DisplayName == "" {
		log.Fatal("You must pass the group display name in DisplayName field.")
	}
	return &apiReplaceGroupRequest{
		Schemas:     []string{defaultGroupSchema},
		DisplayName: group.DisplayName,
		Members:     convertPorcelainToCreateMembersRequest(group.Members),
	}
}

func convertPorcelainToCreateMembersRequest(members []GroupMember) []*apiGroupMemberRequest {
	memberRequestList := []*apiGroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			log.Fatal("You must pass the member value in Value field.")
		} else if member.Display == "" {
			log.Fatal("You must pass the member display in Display field.")
		}
		memberRequestList = append(memberRequestList, &apiGroupMemberRequest{
			Value:   member.Value,
			Display: member.Display,
		})
	}
	return memberRequestList
}

func convertPorcelainToUpdateGroupAddMembers(members []GroupMember) *apiUpdateGroupRequest {
	memberValues := []apiGroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			log.Fatal("You must pass the member value in Value field.")
		} else if member.Display == "" {
			log.Fatal("You must pass the member display in Display field.")
		}
		memberValues = append(memberValues, apiGroupMemberRequest(member))
	}
	return &apiUpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			apiUpdateGroupOperationRequest{
				OP:    "add",
				Path:  "members",
				Value: memberValues,
			},
		},
	}
}

func convertPorcelainToUpdateGroupReplaceMembers(members []GroupMember) *apiUpdateGroupRequest {
	memberValues := []apiGroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			log.Fatal("You must pass the member value in Value field.")
		} else if member.Display == "" {
			log.Fatal("You must pass the member display in Display field.")
		}
		memberValues = append(memberValues, apiGroupMemberRequest(member))
	}
	return &apiUpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			apiUpdateGroupOperationRequest{
				OP:    "replace",
				Path:  "members",
				Value: memberValues,
			},
		},
	}
}

func convertPorcelainToUpdateGroupName(replaceName UpdateGroupReplaceName) *apiUpdateGroupRequest {
	if replaceName.DisplayName == "" {
		log.Fatal("You must pass the group name.")
	}
	return &apiUpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			apiUpdateGroupOperationRequest{
				OP: "replace",
				Value: map[string]string{
					"displayName": replaceName.DisplayName,
				},
			},
		},
	}
}

func convertPorcelainToUpdateGroupRemoveMember(memberID string) *apiUpdateGroupRequest {
	if memberID == "" {
		log.Fatal("You must pass the member id.")
	}
	return &apiUpdateGroupRequest{
		Schemas: []string{defaultGroupSchema},
		Operations: []interface{}{
			&apiUpdateGroupOperationRequest{
				OP:   "remove",
				Path: fmt.Sprintf("members[value eq \"%s\"]", memberID),
			},
		},
	}
}
