package module

import (
	"errors"

	"github.com/strongdm/scimsdk/internal/service"
	"github.com/strongdm/scimsdk/models"
)

const (
	defaultGroupSchema = "urn:ietf:params:scim:schemas:core:2.0:Group"
	defaultPatchSchema = "urn:ietf:params:scim:api:messages:2.0:PatchOp"
	defaultUserSchema  = "urn:ietf:params:scim:schemas:core:2.0:User"
)

func newServiceCreateOptions(body interface{}, url string) *service.CreateOptions {
	return &service.CreateOptions{
		Body:       body,
		BaseAPIURL: url,
	}
}

func newServiceListOptions(opts *models.PaginationOptions, url string) (*service.ListOptions, error) {
	if opts == nil {
		opts = &models.PaginationOptions{}
	}
	if opts.Offset < 0 {
		return nil, errors.New("the pagination offset must be positive")
	} else if opts.PageSize < 0 {
		return nil, errors.New("the pagination page size must be positive")
	}
	return &service.ListOptions{
		PageSize:   opts.PageSize,
		Offset:     opts.Offset,
		Filter:     opts.Filter,
		BaseAPIURL: url,
	}, nil
}

func newServiceFindOptions(id string, url string) (*service.FindOptions, error) {
	if id == "" {
		return nil, errors.New("you must pass the resource id")
	}
	return &service.FindOptions{
		ID:         id,
		BaseAPIURL: url,
	}, nil
}

func newServiceReplaceOptions(id string, body interface{}, url string) (*service.ReplaceOptions, error) {
	if id == "" {
		return nil, errors.New("you must pass the resource id")
	}
	return &service.ReplaceOptions{
		ID:         id,
		Body:       body,
		BaseAPIURL: url,
	}, nil
}

func newServiceUpdateOptions(id string, body interface{}, url string) (*service.UpdateOptions, error) {
	if id == "" {
		return nil, errors.New("you must pass the resource id")
	}
	return &service.UpdateOptions{
		ID:         id,
		Body:       body,
		BaseAPIURL: url,
	}, nil
}

func newServiceDeleteOptions(id string, url string) (*service.DeleteOptions, error) {
	if id == "" {
		return nil, errors.New("you must pass the resource id")
	}
	return &service.DeleteOptions{
		ID:         id,
		BaseAPIURL: url,
	}, nil
}
