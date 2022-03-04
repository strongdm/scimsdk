package scimsdk

import (
	"fmt"
	"log"

	"github.com/strongdm/scimsdk/internal/service"
)

func convertGroupResponseListToPorcelain(groupListResponse []*service.GroupResponse) []*Group {
	groupList := make([]*Group, 0)
	for _, groupResponse := range groupListResponse {
		groupList = append(groupList, convertGroupResponseToPorcelain(groupResponse))
	}
	return groupList
}

func convertGroupResponseToPorcelain(groupResponse *service.GroupResponse) *Group {
	return &Group{
		ID:          groupResponse.ID,
		DisplayName: groupResponse.DisplayName,
		Members:     convertGroupMemberResponseListToPorcelain(groupResponse.Members),
		Meta:        convertGroupMetaResponseToPorcelain(groupResponse.Meta),
	}
}

func convertGroupMemberResponseListToPorcelain(memberListResponse []*service.GroupMemberResponse) []*GroupMember {
	memberList := make([]*GroupMember, 0)
	for _, memberResponse := range memberListResponse {
		memberList = append(memberList, convertGroupMemberResponseToPorcelain(memberResponse))
	}
	return memberList
}

func convertGroupMemberResponseToPorcelain(memberResponse *service.GroupMemberResponse) *GroupMember {
	return &GroupMember{
		Value:   memberResponse.Value,
		Display: memberResponse.Display,
	}
}

func convertGroupMetaResponseToPorcelain(metaResponse *service.GroupMetadataResponse) *GroupMetadata {
	return &GroupMetadata{
		ResourceType: metaResponse.ResourceType,
		Location:     metaResponse.Location,
	}
}

func convertPorcelainToCreateGroupRequest(group *CreateGroupBody) *service.CreateGroupRequest {
	if group.DisplayName == "" {
		log.Fatal("You must pass the group display name in DisplayName field.")
	}
	return &service.CreateGroupRequest{
		Schemas:     []string{defaultGroupSchema},
		DisplayName: group.DisplayName,
		Members:     convertPorcelainToCreateMembersRequest(group.Members),
	}
}

func convertPorcelainToReplaceGroupRequest(group *ReplaceGroupBody) *service.ReplaceGroupRequest {
	if group.DisplayName == "" {
		log.Fatal("You must pass the group display name in DisplayName field.")
	}
	return &service.ReplaceGroupRequest{
		Schemas:     []string{defaultGroupSchema},
		DisplayName: group.DisplayName,
		Members:     convertPorcelainToCreateMembersRequest(group.Members),
	}
}

func convertPorcelainToCreateMembersRequest(members []GroupMember) []*service.GroupMemberRequest {
	memberRequestList := []*service.GroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			log.Fatal("You must pass the member value in Value field.")
		} else if member.Display == "" {
			log.Fatal("You must pass the member display in Display field.")
		}
		memberRequestList = append(memberRequestList, &service.GroupMemberRequest{
			Value:   member.Value,
			Display: member.Display,
		})
	}
	return memberRequestList
}

func convertPorcelainToUpdateGroupAddMembersRequest(members []GroupMember) *service.UpdateGroupRequest {
	memberValues := []service.GroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			log.Fatal("You must pass the member value in Value field.")
		} else if member.Display == "" {
			log.Fatal("You must pass the member display in Display field.")
		}
		memberValues = append(memberValues, service.GroupMemberRequest(member))
	}
	return &service.UpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			service.UpdateGroupOperationRequest{
				OP:    "add",
				Path:  "members",
				Value: memberValues,
			},
		},
	}
}

func convertPorcelainToUpdateGroupReplaceMembersRequest(members []GroupMember) *service.UpdateGroupRequest {
	memberValues := []service.GroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			log.Fatal("You must pass the member value in Value field.")
		} else if member.Display == "" {
			log.Fatal("You must pass the member display in Display field.")
		}
		memberValues = append(memberValues, service.GroupMemberRequest(member))
	}
	return &service.UpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			service.UpdateGroupOperationRequest{
				OP:    "replace",
				Path:  "members",
				Value: memberValues,
			},
		},
	}
}

func convertPorcelainToUpdateGroupNameRequest(replaceName UpdateGroupReplaceName) *service.UpdateGroupRequest {
	if replaceName.DisplayName == "" {
		log.Fatal("You must pass the group name.")
	}
	return &service.UpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			service.UpdateGroupOperationRequest{
				OP: "replace",
				Value: map[string]string{
					"displayName": replaceName.DisplayName,
				},
			},
		},
	}
}

func convertPorcelainToUpdateGroupRemoveMemberRequest(memberID string) *service.UpdateGroupRequest {
	if memberID == "" {
		log.Fatal("You must pass the member id.")
	}
	return &service.UpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			&service.UpdateGroupOperationRequest{
				OP:   "remove",
				Path: fmt.Sprintf("members[value eq \"%s\"]", memberID),
			},
		},
	}
}
