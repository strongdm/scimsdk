package scimsdk

import (
	"log"

	"github.com/strongdm/scimsdk/internal/service"
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

type PaginationOptions struct {
	PageSize int
	Offset   int
	Filter   string
}

func newServiceListOptions(opts *PaginationOptions, url string) *service.ListOptions {
	if opts == nil {
		opts = &PaginationOptions{}
	}
	if opts.Offset < 0 {
		log.Fatal("The pagination offset must be positive")
	} else if opts.PageSize < 0 {
		log.Fatal("The pagination page size must be positive")
	}
	return &service.ListOptions{
		PageSize:   opts.PageSize,
		Offset:     opts.Offset,
		Filter:     opts.Filter,
		BaseAPIURL: url,
	}
}

func newServiceFindOptions(id string, url string) *service.FindOptions {
	if id == "" {
		log.Fatal("You must pass the resource id")
	}
	return &service.FindOptions{
		ID:         id,
		BaseAPIURL: url,
	}
}

func newServiceReplaceOptions(id string, body interface{}, url string) *service.ReplaceOptions {
	if id == "" {
		log.Fatal("You must pass the resource id")
	}
	return &service.ReplaceOptions{
		ID:         id,
		Body:       body,
		BaseAPIURL: url,
	}
}

func newServiceUpdateOptions(id string, body interface{}, url string) *service.UpdateOptions {
	if id == "" {
		log.Fatal("You must pass the resource id")
	}
	return &service.UpdateOptions{
		ID:         id,
		Body:       body,
		BaseAPIURL: url,
	}
}

func newServiceDeleteOptions(id string, url string) *service.DeleteOptions {
	if id == "" {
		log.Fatal("You must pass the resource id")
	}
	return &service.DeleteOptions{
		ID:         id,
		BaseAPIURL: url,
	}
}
