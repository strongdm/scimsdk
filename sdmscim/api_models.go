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

type apiCreateUserRequest struct {
	Schemas  []string           `json:"schemas"`
	UserName string             `json:"userName"`
	Name     apiUserNameRequest `json:"name"`
	Active   bool               `json:"active"`
}

type apiReplaceUserRequest struct {
	ID       string             `json:"id"`
	Schemas  []string           `json:"schemas"`
	UserName string             `json:"userName"`
	Name     apiUserNameRequest `json:"name"`
	Active   bool               `json:"active"`
}

type apiUserNameRequest struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type apiUpdateUserRequest struct {
	Schemas    []string                    `json:"schemas"`
	Operations []apiUpdateOperationRequest `json:"Operations"`
}

type apiUpdateOperationRequest struct {
	OP    string                         `json:"op"`
	Value apiUpdateOperationValueRequest `json:"value"`
}

type apiUpdateOperationValueRequest struct {
	Active bool `json:"active"`
}
