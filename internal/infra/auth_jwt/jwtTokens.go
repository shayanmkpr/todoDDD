package auth_jwt

import (
	"time"

	"todoDB/internal/domain/auth"

	"github.com/golang-jwt/jwt"
)

type authRepository struct{}

func NewAuthRepository() auth.AuthenticationRepo {
	return &authRepository{}
}

func (auth_ptr *authRepository) GenerateAccessToken(secret, userID string) (string, error) {
	now := time.Now()

	claims := &auth.Claims{
		UserID:    userID,
		IssuedAt:  now,
		ExpiresAt: now.Add(auth.AccressTokenDuration),
	}

	return GenerateToken(secret, userID, claims)
}

func (auth_ptr *authRepository) GenerateRefreshToken(secret, userID string) (string, error) {
	now := time.Now()

	claims := &auth.Claims{
		UserID:    userID,
		IssuedAt:  now,
		ExpiresAt: now.Add(auth.RefreshTokenDuration),
	}

	return GenerateToken(secret, userID, claims)
}

func (auth_ptr *authRepository) ValidateAccessToken(secret, tokenStr string) (bool, error) {
	_, err := ParseToken(secret, tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (auth_ptr *authRepository) ValidateRefreshToken(secret, tokenStr string) (bool, error) {
	return false, nil
}

func GenerateToken(secret, userID string, claims *auth.Claims) (string, error) {
	jwtClaims := jwt.MapClaims{
		"user_id": claims.UserID,
		"iat":     claims.IssuedAt.Unix(),
		"exp":     claims.ExpiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenStr string) (*auth.Claims, error) {
	jwtClaims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(tokenStr, jwtClaims,
		func(t *jwt.Token) (interface{}, error) { // key function
			return []byte(secret), nil
		},
	)

	if err != nil || !t.Valid {
		return nil, err
	}

	m := t.Claims.(jwt.MapClaims)

	return &auth.Claims{
		UserID:    m["user_id"].(string),
		IssuedAt:  time.Unix(int64(m["iat"].(float64)), 0),
		ExpiresAt: time.Unix(int64(m["exp"].(float64)), 0),
	}, nil
}
