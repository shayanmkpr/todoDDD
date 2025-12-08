package auth

import (
	"time"
)

type Claims struct {
	UserName  string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

type AccessToken struct {
	Value  string
	Claims Claims
}

type RefreshToken struct { // gormable
	Value  string
	Claims Claims
}

const (
	RefreshTokenDuration = 3 * 24 * time.Hour
	AccressTokenDuration = 15 * time.Minute
)
