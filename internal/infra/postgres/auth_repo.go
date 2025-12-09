package postgres

import (
	"context"
	"todoDB/internal/domain/auth"

	"gorm.io/gorm"
)

type refreshTokenRepo struct {
	db *gorm.DB
}

func NewAuthPostgresRepository(db *gorm.DB) auth.RefreshTokenRepo {
	return &refreshTokenRepo{db: db}
}

func (t *refreshTokenRepo) StoreRefreshToken(ctx context.Context, token *auth.RefreshToken) error {
	return t.db.WithContext(ctx).Create(token).Error
}

func (t *refreshTokenRepo) GetRefreshToken(ctx context.Context, tokenValue string) (*auth.RefreshToken, error) {
	var theToken auth.RefreshToken
	err := t.db.WithContext(ctx).Where("value = ?", tokenValue).First(&theToken).Error
	if err != nil {
		return nil, err
	} else {
		return &theToken, nil
	}
}

func (t *refreshTokenRepo) DeleteRefershToken(ctx context.Context, userName string) error {
	return t.db.WithContext(ctx).Where("user_name = ?", userName).Delete(&auth.RefreshToken{}).Error
}
