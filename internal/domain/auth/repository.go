package auth

import "context"

type RefreshTokenRepo interface { // for postgres
	StoreRefreshToken(ctx context.Context, token RefreshToken) error
	GetRefreshToken(ctx context.Context, userName string) (string, error)
	DeleteRefershToken(ctx context.Context, userName string) error
}

type AuthenticationRepo interface {
	GenerateAccessToken(secret, userID string) (string, error)
	GenerateRefreshToken(secret, userID string) (string, error)
	ValidateAccessToken(secret, toeknStr string) (bool, error)
	ValidateRefreshToken(secret, toeknStr string) (bool, error)
}
