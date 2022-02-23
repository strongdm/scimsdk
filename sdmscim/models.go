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
