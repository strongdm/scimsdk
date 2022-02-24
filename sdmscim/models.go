package sdmscim

type User struct {
	ID          string
	Active      bool
	DisplayName string
	Emails      []UserEmail
	Groups      []interface{}
	Name        *UserName
	UserName    string
	UserType    string
}

type UserEmail struct {
	Primary bool
	Value   string
}

type UserName struct {
	FamilyName string
	Formatted  string
	GivenName  string
}

type CreateUserBody struct {
	UserName   string
	GivenName  string
	FamilyName string
	Active     bool
}

type ReplaceUserBody struct {
	ID         string
	UserName   string
	GivenName  string
	FamilyName string
	Active     bool
}
