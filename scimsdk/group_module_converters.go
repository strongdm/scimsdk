package scimsdk

import (
	"errors"
	"fmt"

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

func convertPorcelainToCreateGroupRequest(group *CreateGroupBody) (*service.CreateGroupRequest, error) {
	if group.DisplayName == "" {
		return nil, errors.New("you must pass the group display name in DisplayName field")
	}
	members, err := convertPorcelainToCreateMembersRequest(group.Members)
	if err != nil {
		return nil, err
	}
	return &service.CreateGroupRequest{
		Schemas:     []string{defaultGroupSchema},
		DisplayName: group.DisplayName,
		Members:     members,
	}, nil
}

func convertPorcelainToReplaceGroupRequest(group *ReplaceGroupBody) (*service.ReplaceGroupRequest, error) {
	if group.DisplayName == "" {
		return nil, errors.New("you must pass the group display name in DisplayName field")
	}
	members, err := convertPorcelainToCreateMembersRequest(group.Members)
	if err != nil {
		return nil, err
	}
	return &service.ReplaceGroupRequest{
		Schemas:     []string{defaultGroupSchema},
		DisplayName: group.DisplayName,
		Members:     members,
	}, nil
}

func convertPorcelainToCreateMembersRequest(members []GroupMember) ([]*service.GroupMemberRequest, error) {
	memberRequestList := []*service.GroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			return nil, errors.New("you must pass the member value in Value field")
		} else if member.Display == "" {
			return nil, errors.New("you must pass the member display in Display field")
		}
		memberRequestList = append(memberRequestList, &service.GroupMemberRequest{
			Value:   member.Value,
			Display: member.Display,
		})
	}
	return memberRequestList, nil
}

func convertPorcelainToUpdateGroupAddMembersRequest(members []GroupMember) (*service.UpdateGroupRequest, error) {
	memberValues := []service.GroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			return nil, errors.New("you must pass the member value in Value field")
		} else if member.Display == "" {
			return nil, errors.New("you must pass the member display in Display field")
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
	}, nil
}

func convertPorcelainToUpdateGroupReplaceMembersRequest(members []GroupMember) (*service.UpdateGroupRequest, error) {
	memberValues := []service.GroupMemberRequest{}
	for _, member := range members {
		if member.Value == "" {
			return nil, errors.New("you must pass the member value in Value field")
		} else if member.Display == "" {
			return nil, errors.New("you must pass the member display in Display field")
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
	}, nil
}

func convertPorcelainToUpdateGroupNameRequest(replaceName UpdateGroupReplaceName) (*service.UpdateGroupRequest, error) {
	if replaceName.DisplayName == "" {
		return nil, errors.New("you must pass the group name in DisplayName field")
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
	}, nil
}

func convertPorcelainToUpdateGroupRemoveMemberRequest(memberID string) (*service.UpdateGroupRequest, error) {
	if memberID == "" {
		return nil, errors.New("you must pass the member id in memberID field")
	}
	return &service.UpdateGroupRequest{
		Schemas: []string{defaultPatchSchema},
		Operations: []interface{}{
			&service.UpdateGroupOperationRequest{
				OP:   "remove",
				Path: fmt.Sprintf("members[value eq \"%s\"]", memberID),
			},
		},
	}, nil
}
