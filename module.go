package scimsdk

import (
	"context"

	"github.com/strongdm/scimsdk/models"
)

type UserModule interface {
	Create(context.Context, models.CreateUser) (*models.User, error)
	List(context.Context, *models.PaginationOptions) models.Iterator[models.User]
	Find(context.Context, string) (*models.User, error)
	Replace(context.Context, string, models.ReplaceUser) (*models.User, error)
	Update(context.Context, string, models.UpdateUser) (bool, error)
	Delete(context.Context, string) (bool, error)
}

type GroupModule interface {
	Create(context.Context, models.CreateGroupBody) (*models.Group, error)
	List(context.Context, *models.PaginationOptions) models.Iterator[models.Group]
	Find(context.Context, string) (*models.Group, error)
	Replace(context.Context, string, models.ReplaceGroupBody) (*models.Group, error)
	UpdateAddMembers(context.Context, string, []models.GroupMember) (bool, error)
	UpdateReplaceMembers(context.Context, string, []models.GroupMember) (bool, error)
	UpdateReplaceName(context.Context, string, models.UpdateGroupReplaceName) (bool, error)
	UpdateRemoveMemberByID(context.Context, string, string) (bool, error)
	Delete(context.Context, string) (bool, error)
}
