package auth

import (
	"time"

	"gorm.io/gorm"
)

type Claims struct {
	UserName  string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

type RefreshToken struct {
	gorm.Model
	Value     string    `gorm:"unique;not null"`
	UserName  string    `gorm:"not null"`
	IssuedAt  time.Time `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}

const (
	RefreshTokenDuration = 3 * 24 * time.Hour
	AccressTokenDuration = 15 * time.Minute
)
