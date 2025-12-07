package user

import "context"

type UserRepository interface {
	DeleteUser(ctx context.Context, userName string) error
	GetByName(ctx context.Context, userName string) (*User, error)
	CheckByName(ctx context.Context, userName string) (bool, error)
	SaveUser(ctx context.Context, newUser *User) error
}
