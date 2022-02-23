package sdmscim

type serviceCreateOptions struct {
	Body       interface{}
	BaseAPIURL string
}

// serviceListOptions is implemente using as reference the strongDM SCIM docs
// (https://www.strongdm.com/docs/architecture/scim-spec/users/list) using:
// - count -> PageSize (default value is 5)
// - startIndex -> offset (default value is 1)
// - filter -> filter
type serviceListOptions struct {
	// PageSize defines the count of the users by page
	PageSize int
	// Offset defines the page offset referencing to the page - relative to the PageSize
	Offset int
	// Filter defines the query filter used in strongDM
	Filter     string
	BaseAPIURL string
}

type serviceFindOptions struct {
	ID         string
	BaseAPIURL string
}

type serviceDeleteOptions struct {
	ID         string
	BaseAPIURL string
}

type serviceReplaceOptions struct {
	ID         string
	Body       interface{}
	BaseAPIURL string
}
