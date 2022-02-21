package api

type APIUserPageResponse struct {
	Resources    []APIUserResponse `json:"Resources"`
	ItemsPerPage int               `json:"itemsPerPage"`
	Schemas      []string          `json:"schemas"`
	StartIndex   int               `json:"startIndex"`
	TotalResults int               `json:"totalResults"`
}

type APIUserResponse struct {
	ID          string                 `json:"id"`
	Active      bool                   `json:"active"`
	DisplayName string                 `json:"displayName"`
	Emails      []APIUserEmailResponse `json:"emails"`
	Groups      []interface{}          `json:"groups"`
	Name        APIUserNameResponse    `json:"name"`
	Schemas     []string               `json:"schemas"`
	UserName    string                 `json:"userName"`
	UserType    string                 `json:"userType"`
}

type APIUserEmailResponse struct {
	Primary bool   `json:"primary"`
	Value   string `json:"value"`
}

type APIUserNameResponse struct {
	FamilyName string `json:"familyName"`
	Formatted  string `json:"formatted"`
	GivenName  string `json:"givenName"`
}

type APIGroupPageResponse struct {
	Resources    []APIGroupResponse `json:"Resources"`
	ItemsPerPage int                `json:"itemsPerPage"`
	Schemas      []string           `json:"schemas"`
	StartIndex   int                `json:"startIndex"`
	TotalResults int                `json:"totalResults"`
}

type APIGroupResponse struct {
	ID          string
	DisplayName string
	Members     []interface{}
	Meta        APIGroupMetaResponse
	Schemas     []string
}

type APIGroupMetaResponse struct {
	ResourceType string
	Location     string
}
