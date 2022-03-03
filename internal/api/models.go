package api

type CreateOptions struct {
	Body       interface{}
	BaseAPIURL string
}

// ListOptions is implemented using the strongDM SCIM docs as reference
// (https://www.strongdm.com/docs/architecture/scim-spec/users/list) using:
// - count -> PageSize (default value is 5)
// - startIndex -> offset (default value is 1)
// - filter -> filter
type ListOptions struct {
	// PageSize defines the resource count by page
	PageSize int
	// Offset defines the page offset referencing to the page - relative to the PageSize
	Offset int
	// Filter defines the query filter used in strongDM
	Filter     string
	BaseAPIURL string
}

type FindOptions struct {
	ID         string
	BaseAPIURL string
}

type ReplaceOptions struct {
	ID         string
	Body       interface{}
	BaseAPIURL string
}

type UpdateOptions struct {
	ID         string
	Body       interface{}
	BaseAPIURL string
}

type DeleteOptions struct {
	ID         string
	BaseAPIURL string
}

func NewCreateOptions(body interface{}, baseAPIURL string) *CreateOptions {
	return &CreateOptions{body, baseAPIURL}
}

func NewListOptions(pageSize, offset int, filter, baseAPIURL string) *ListOptions {
	return &ListOptions{pageSize, offset, filter, baseAPIURL}
}

func NewFindOptions(id, baseAPIURL string) *FindOptions {
	return &FindOptions{id, baseAPIURL}
}

func NewReplaceOptions(id string, body interface{}, baseAPIURL string) *ReplaceOptions {
	return &ReplaceOptions{id, body, baseAPIURL}
}

func NewUpdateOptions(id string, body interface{}, baseAPIURL string) *UpdateOptions {
	return &UpdateOptions{id, body, baseAPIURL}
}

func NewDeleteOptions(id, baseAPIURL string) *DeleteOptions {
	return &DeleteOptions{id, baseAPIURL}
}
