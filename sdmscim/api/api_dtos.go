package api

type APIUserPageResponseDTO struct {
	Resources    []APIUserResponseDTO `json:"Resources"`
	ItemsPerPage int                  `json:"itemsPerPage"`
	Schemas      []string             `json:"schemas"`
	StartIndex   int                  `json:"startIndex"`
	TotalResults int                  `json:"totalResults"`
}

type APIUserResponseDTO struct {
	ID          string                    `json:"id"`
	Active      bool                      `json:"active"`
	DisplayName string                    `json:"displayName"`
	Emails      []APIUserEmailResponseDTO `json:"emails"`
	Groups      []interface{}             `json:"groups"`
	Name        APIUserNameResponseDTO    `json:"name"`
	Schemas     []string                  `json:"schemas"`
	UserName    string                    `json:"userName"`
	UserType    string                    `json:"userType"`
}

type APIUserEmailResponseDTO struct {
	Primary bool   `json:"primary"`
	Value   string `json:"value"`
}

type APIUserNameResponseDTO struct {
	FamilyName string `json:"familyName"`
	Formatted  string `json:"formatted"`
	GivenName  string `json:"givenName"`
}

type APIGroupPageResponseDTO struct {
	Resources    []APIGroupResponseDTO `json:"Resources"`
	ItemsPerPage int                   `json:"itemsPerPage"`
	Schemas      []string              `json:"schemas"`
	StartIndex   int                   `json:"startIndex"`
	TotalResults int                   `json:"totalResults"`
}

type APIGroupResponseDTO struct {
	ID          string
	DisplayName string
	Members     []interface{}
	Meta        APIGroupMetaResponseDTO
	Schemas     []string
}

type APIGroupMetaResponseDTO struct {
	ResourceType string
	Location     string
}
