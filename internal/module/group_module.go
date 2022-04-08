package module

import (
	"context"

	"github.com/strongdm/scimsdk/internal/service"
	"github.com/strongdm/scimsdk/models"
)

type groupModuleImpl struct {
	service     service.GroupService
	providedURL string
}

func NewGroupModule(service service.GroupService, providedURL string) *groupModuleImpl {
	return &groupModuleImpl{service, providedURL}
}

func (module *groupModuleImpl) Create(ctx context.Context, group models.CreateGroupBody) (*models.Group, error) {
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

func (module *groupModuleImpl) List(ctx context.Context, paginationOptions *models.PaginationOptions) models.Iterator[models.Group] {
	return newIterator(module.iteratorMiddleware(ctx), paginationOptions)
}

func (module *groupModuleImpl) Find(ctx context.Context, id string) (*models.Group, error) {
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

func (module *groupModuleImpl) Replace(ctx context.Context, id string, group models.ReplaceGroupBody) (*models.Group, error) {
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

func (module *groupModuleImpl) UpdateAddMembers(ctx context.Context, id string, members []models.GroupMember) (bool, error) {
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

func (module *groupModuleImpl) UpdateReplaceMembers(ctx context.Context, id string, members []models.GroupMember) (bool, error) {
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

func (module *groupModuleImpl) UpdateReplaceName(ctx context.Context, id string, replaceName models.UpdateGroupReplaceName) (bool, error) {
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

func (module *groupModuleImpl) UpdateRemoveMemberByID(ctx context.Context, id string, memberID string) (bool, error) {
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

func (module *groupModuleImpl) Delete(ctx context.Context, id string) (bool, error) {
	opts, err := newServiceDeleteOptions(id, module.providedURL)
	if err != nil {
		return false, err
	}
	return module.service.Delete(ctx, opts)
}

func (module *groupModuleImpl) iteratorMiddleware(ctx context.Context) iteratorFetchFunc[models.Group] {
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
