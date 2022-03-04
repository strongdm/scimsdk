package service

type UserPageResponse struct {
	Resources    []UserResponse `json:"Resources"`
	ItemsPerPage int            `json:"itemsPerPage"`
	Schemas      []string       `json:"schemas"`
	StartIndex   int            `json:"startIndex"`
	TotalResults int            `json:"totalResults"`
}

type UserResponse struct {
	ID          string              `json:"id"`
	Active      bool                `json:"active"`
	DisplayName string              `json:"displayName"`
	Emails      []UserEmailResponse `json:"emails"`
	Groups      []interface{}       `json:"groups"`
	Name        UserNameResponse    `json:"name"`
	Schemas     []string            `json:"schemas"`
	UserName    string              `json:"userName"`
	UserType    string              `json:"userType"`
}

type UserEmailResponse struct {
	Primary bool   `json:"primary"`
	Value   string `json:"value"`
}

type UserNameResponse struct {
	FamilyName string `json:"familyName"`
	Formatted  string `json:"formatted"`
	GivenName  string `json:"givenName"`
}

type CreateUserRequest struct {
	Schemas  []string        `json:"schemas"`
	UserName string          `json:"userName"`
	Name     UserNameRequest `json:"name"`
	Active   bool            `json:"active"`
}

type ReplaceUserRequest struct {
	ID       string          `json:"id"`
	Schemas  []string        `json:"schemas"`
	UserName string          `json:"userName"`
	Name     UserNameRequest `json:"name"`
	Active   bool            `json:"active"`
}

type UserNameRequest struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type UpdateUserRequest struct {
	Schemas    []string                     `json:"schemas"`
	Operations []UpdateUserOperationRequest `json:"Operations"`
}

type UpdateUserOperationRequest struct {
	OP    string                          `json:"op"`
	Value UpdateUserOperationValueRequest `json:"value"`
}

type UpdateUserOperationValueRequest struct {
	Active bool `json:"active"`
}
