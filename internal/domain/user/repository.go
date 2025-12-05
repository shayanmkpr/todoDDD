package user

import "context"

type UserRepository interface {
	UpdateUserName(ctx context.Context, newUserName string) error
	UpdateUserPass(ctx context.Context, newUserPass string) error
	DeleteUser(ctx context.Context, userName string) error
	GetByNme(ctx context.Context, userName string) (*User, error)
	CheckByName(ctx context.Context, userName string) (bool, error)
	CreateUser(ctx context.Context, newUser *User) error
}
