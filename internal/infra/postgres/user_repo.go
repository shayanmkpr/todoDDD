package postgres

import (
	"context"

	"todoDB/internal/domain/user"

	"gorm.io/gorm"
)

type userRepository struct { // local to db repository
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) SaveUser(ctx context.Context, newUser *user.User) error {
	return r.db.WithContext(ctx).Create(newUser).Error
}

func (r *userRepository) GetByName(ctx context.Context, userName string) (*user.User, error) {
	var theUser user.User
	err := r.db.WithContext(ctx).Where("user_name = ?", userName).First(&theUser).Error
	if err != nil {
		return nil, err
	} else {
		return &theUser, nil
	}
}

func (r *userRepository) CheckByName(ctx context.Context, userName string) (bool, error) {
	var theUser user.User
	err := r.db.WithContext(ctx).Where("user_name = ?", userName).Find(&theUser).Error
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (r *userRepository) DeleteUser(ctx context.Context, userName string) error {
	return r.db.WithContext(ctx).Where("user_name = ?", userName).Delete(&user.User{}).Error
}
