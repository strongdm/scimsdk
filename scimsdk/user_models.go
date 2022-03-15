package scimsdk

type User struct {
	ID          string
	Active      bool
	DisplayName string
	Emails      []UserEmail
	Groups      []UserGroupReference
	Name        *UserName
	UserName    string
	UserType    string
}

type UserEmail struct {
	Primary bool
	Value   string
}

type UserGroupReference struct {
	Value string
	Ref   string
}

type UserName struct {
	FamilyName string
	Formatted  string
	GivenName  string
}

type CreateUser struct {
	UserName   string
	GivenName  string
	FamilyName string
	Active     bool
}

type ReplaceUser CreateUser

type UpdateUser struct {
	Active bool
}
