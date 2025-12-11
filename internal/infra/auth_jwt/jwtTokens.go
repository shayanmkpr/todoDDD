package auth_jwt

import (
	"context"
	"time"

	"todoDB/internal/domain/auth"

	"github.com/golang-jwt/jwt"
)

type authRepository struct{} // the only reason that Im doing this is cause Im trying DDD. This is stupid.

func NewAuthRepository() auth.AuthenticationRepo {
	return &authRepository{}
}

func (auth_ptr *authRepository) GenerateAccessToken(ctx context.Context, secret, userName string) (string, error) {
	now := time.Now()

	claims := &auth.Claims{
		UserName:  userName,
		IssuedAt:  now,
		ExpiresAt: now.Add(auth.AccressTokenDuration),
	}

	return GenerateToken(secret, userName, claims)
}

func (auth_ptr *authRepository) GenerateRefreshToken(ctx context.Context, secret, userName string) (auth.RefreshToken, error) {
	now := time.Now()

	claims := &auth.Claims{
		UserName:  userName,
		IssuedAt:  now,
		ExpiresAt: now.Add(auth.RefreshTokenDuration),
	}

	tokenString, err := GenerateToken(secret, userName, claims)
	if err != nil {
	}

	return
}

func GenerateToken(secret, userName string, claims *auth.Claims) (string, error) {
	jwtClaims := jwt.MapClaims{
		"user_id": claims.UserName,
		"iat":     claims.IssuedAt.Unix(),
		"exp":     claims.ExpiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(secret))
}

func (auth_ptr *authRepository) ParseToken(ctx context.Context, secret, tokenStr string) (*auth.Claims, error) {
	jwtClaims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(tokenStr, jwtClaims,
		func(t *jwt.Token) (interface{}, error) { // key function. this guy checks the secret?
			return []byte(secret), nil
		},
	)

	if err != nil || !t.Valid {
		return nil, err
	}

	m := t.Claims.(jwt.MapClaims)

	return &auth.Claims{
		UserName:  m["user_id"].(string),
		IssuedAt:  time.Unix(int64(m["iat"].(float64)), 0),
		ExpiresAt: time.Unix(int64(m["exp"].(float64)), 0),
	}, nil
}
