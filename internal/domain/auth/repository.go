package auth

import "context"

type RefreshTokenRepo interface { // for postgres
	StoreRefreshToken(ctx context.Context, token *RefreshToken) error
	GetRefreshToken(ctx context.Context, tokenValue string) (*RefreshToken, error)
	DeleteRefershToken(ctx context.Context, tokenValue string) error
}

type AuthenticationRepo interface {
	GenerateAccessToken(ctx context.Context, secret, userName string) (string, error)
	GenerateRefreshToken(ctx context.Context, secret, userName string) (*RefreshToken, error)
	ParseToken(ctx context.Context, secret, tokenStr string) (*Claims, error)
}
