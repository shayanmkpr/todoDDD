package authenticatoin_jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID    string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

var (
	AccessTokenSecret  string = os.Getenv("ACCESS_TOKEN_SECRET")
	RefreshTokenSecret string = os.Getenv("REFRESH_TOKEN_SECRET")
)

func GenerateAccessToken(secret, userID string) (string, error) {
	now := time.Now()

	claims := &Claims{
		UserID:    userID,
		IssuedAt:  now,
		ExpiresAt: now.Add(15 * time.Minute), // typical access token lifetime
	}

	return GenerateToken(secret, userID, claims)
}

func GenerateRefreshToken(secret, userID string) (string, error) {
	now := time.Now()

	claims := &Claims{
		UserID:    userID,
		IssuedAt:  now,
		ExpiresAt: now.Add(30 * 24 * time.Hour), // 30 days
	}

	return GenerateToken(secret, userID, claims)
}

func GenerateToken(secret, userID string, claims *Claims) (string, error) {
	jwtClaims := jwt.MapClaims{
		"user_id": claims.UserID,
		"iat":     claims.IssuedAt.Unix(),
		"exp":     claims.ExpiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenStr string) (*Claims, error) {
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

	return &Claims{
		UserID:    m["user_id"].(string),
		IssuedAt:  time.Unix(int64(m["iat"].(float64)), 0),
		ExpiresAt: time.Unix(int64(m["exp"].(float64)), 0),
	}, nil
}
