package auth

import (
	"time"
)

type Claims struct {
	UserName  string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

type RefreshToken struct {
	Value     string
	UserName  string
	ExpiresAt time.Time
}

const (
	RefreshTokenDuration = 3 * 24 * time.Hour
	AccressTokenDuration = 15 * time.Minute
)
