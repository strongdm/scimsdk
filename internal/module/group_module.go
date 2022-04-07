package module

import (
	"context"

	"github.com/strongdm/scimsdk/internal/service"
	"github.com/strongdm/scimsdk/models"
)

type GroupModule interface {
	Create(ctx context.Context, group models.CreateGroupBody) (*models.Group, error)
	List(ctx context.Context, paginationOptions *models.PaginationOptions) GroupIterator
	Find(ctx context.Context, id string) (*models.Group, error)
	Replace(ctx context.Context, id string, group models.ReplaceGroupBody) (*models.Group, error)
	UpdateAddMembers(ctx context.Context, id string, members []models.GroupMember) (bool, error)
	UpdateReplaceMembers(ctx context.Context, id string, members []models.GroupMember) (bool, error)
	UpdateReplaceName(ctx context.Context, id string, replaceName models.UpdateGroupReplaceName) (bool, error)
	UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type GroupModuleImpl struct {
	service     service.GroupService
	providedURL string
}

func NewGroupModule(service service.GroupService, providedURL string) *GroupModuleImpl {
	return &GroupModuleImpl{service, providedURL}
}

func (module *GroupModuleImpl) Create(ctx context.Context, group models.CreateGroupBody) (*models.Group, error) {
	body, err := convertPorcelainToCreateGroupRequest(&group)
	if err != nil {
		return nil, err
	}
	opts := newServiceCreateOptions(body, module.providedURL)
	response, err := module.service.Create(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModuleImpl) List(ctx context.Context, paginationOptions *models.PaginationOptions) GroupIterator {
	return newGroupsIterator(module.iteratorMiddleware(ctx), paginationOptions)
}

func (module *GroupModuleImpl) Find(ctx context.Context, id string) (*models.Group, error) {
	opts, err := newServiceFindOptions(id, module.providedURL)
	if err != nil {
		return nil, err
	}
	response, err := module.service.Find(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModuleImpl) Replace(ctx context.Context, id string, group models.ReplaceGroupBody) (*models.Group, error) {
	body, err := convertPorcelainToReplaceGroupRequest(&group)
	if err != nil {
		return nil, err
	}
	opts, err := newServiceReplaceOptions(id, body, module.providedURL)
	if err != nil {
		return nil, err
	}
	response, err := module.service.Replace(ctx, opts)
	if err != nil {
		return nil, err
	}
	return convertGroupResponseToPorcelain(response), nil
}

func (module *GroupModuleImpl) UpdateAddMembers(ctx context.Context, id string, members []models.GroupMember) (bool, error) {
	body, err := convertPorcelainToUpdateGroupAddMembersRequest(members)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModuleImpl) UpdateReplaceMembers(ctx context.Context, id string, members []models.GroupMember) (bool, error) {
	body, err := convertPorcelainToUpdateGroupReplaceMembersRequest(members)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModuleImpl) UpdateReplaceName(ctx context.Context, id string, replaceName models.UpdateGroupReplaceName) (bool, error) {
	body, err := convertPorcelainToUpdateGroupNameRequest(replaceName)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModuleImpl) UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error) {
	body, err := convertPorcelainToUpdateGroupRemoveMemberRequest(memberID)
	if err != nil {
		return false, err
	}
	opts, err := newServiceUpdateOptions(id, body, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Update(ctx, opts)
}

func (module *GroupModuleImpl) Delete(ctx context.Context, id string) (bool, error) {
	opts, err := newServiceDeleteOptions(id, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Delete(ctx, opts)
}

func (module *GroupModuleImpl) iteratorMiddleware(ctx context.Context) listGroupsOperationFunc {
	return func(opts *models.PaginationOptions) ([]*models.Group, bool, error) {
		listOpts, err := newServiceListOptions(opts, module.providedURL)
		if err != nil {
			return nil, false, err
		}
		response, haveNextPage, err := module.service.List(ctx, listOpts)
		if err != nil {
			return nil, false, err
		}
		groups := convertGroupResponseListToPorcelain(response)
		return groups, haveNextPage, nil
	}
}
