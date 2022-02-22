package sdmscim

type apiUserPageResponse struct {
	Resources    []apiUserResponse `json:"Resources"`
	ItemsPerPage int               `json:"itemsPerPage"`
	Schemas      []string          `json:"schemas"`
	StartIndex   int               `json:"startIndex"`
	TotalResults int               `json:"totalResults"`
}

type apiUserResponse struct {
	ID          string                 `json:"id"`
	Active      bool                   `json:"active"`
	DisplayName string                 `json:"displayName"`
	Emails      []apiUserEmailResponse `json:"emails"`
	Groups      []interface{}          `json:"groups"`
	Name        apiUserNameResponse    `json:"name"`
	Schemas     []string               `json:"schemas"`
	UserName    string                 `json:"userName"`
	UserType    string                 `json:"userType"`
}

type apiUserEmailResponse struct {
	Primary bool   `json:"primary"`
	Value   string `json:"value"`
}

type apiUserNameResponse struct {
	FamilyName string `json:"familyName"`
	Formatted  string `json:"formatted"`
	GivenName  string `json:"givenName"`
}

// apiListOptions is implemente using as reference the strongDM SCIM docs
// (https://www.strongdm.com/docs/architecture/scim-spec/users/list) using:
// - count -> PageSize (default value is 5)
// - startIndex -> offset (default value is 1)
// - filter -> filter
type apiListOptions struct {
	// PageSize defines the count of the users by page
	PageSize int
	// Offset defines the page offset referencing to the page - relative to the PageSize
	Offset int
	// Filter defines the query filter used in strongDM
	Filter string

	BaseAPIURL string
}
