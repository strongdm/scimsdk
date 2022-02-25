package sdmscim

type serviceCreateOptions struct {
	Body       interface{}
	BaseAPIURL string
}

// serviceListOptions is implemented using the strongDM SCIM docs as reference
// (https://www.strongdm.com/docs/architecture/scim-spec/users/list) using:
// - count -> PageSize (default value is 5)
// - startIndex -> offset (default value is 1)
// - filter -> filter
type serviceListOptions struct {
	// PageSize defines the resource count by page
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

type serviceReplaceOptions struct {
	ID         string
	Body       interface{}
	BaseAPIURL string
}

type serviceUpdateOptions struct {
	ID         string
	Body       interface{}
	BaseAPIURL string
}

type serviceDeleteOptions struct {
	ID         string
	BaseAPIURL string
}
