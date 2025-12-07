package auth

import "context"

type RefreshTokenRepo interface { // for postgres
	StoreRefreshToken(ctx context.Context, token RefreshToken) error
	GetRefreshToken(ctx context.Context, userName string) (string, error)
	DeleteRefershToken(ctx context.Context, userName string) error
}

type AuthenticationRepo interface {
	GenerateAccessToken(ctx context.Context, secret, userID string) (string, error)
	GenerateRefreshToken(ctx context.Context, secret, userID string) (string, error)
	ValidateToken(ctx context.Context, secret, toeknStr string) (bool, error)
}
