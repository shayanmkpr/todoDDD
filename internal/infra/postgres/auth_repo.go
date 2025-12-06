package authentication_postgres

import (
	"context"
	jwt "todoDB/internal/infra/auth"
)

type RefreshTokenRepo interface {
	StoreRefreshToken(ctx context.Context, token RefreshToken) error
	GetRefreshToken(ctx context.Context, userName string) (string, error)
	DeleteRefershToken(ctx context.Context, userName string) error
}

type AccessToken struct {
	Value  string
	Claims jwt.Claims
}

type RefreshToken struct {
	Value  string
	Claims jwt.Claims
}
