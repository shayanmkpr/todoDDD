package user

type UserRepository interface {
	UpdateUserName(newUserName string) error
	UpdateUserPass(newUserPass string) error
	DeleteUser(userName string) error
	GetByNme(userName string) (*User, error)
}
