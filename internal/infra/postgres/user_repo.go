package postgres

import (
	"context"
	"fmt"

	"todoDB/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct { // local to db repository
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{db: db.Model(&user.User{})}
}

func (r *userRepository) SaveUser(ctx context.Context, newUser *user.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedUser := &user.User{
		UserName: newUser.UserName,
		Pass:     string(hashedPass),
	}
	return r.db.WithContext(ctx).Create(hashedUser).Error
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
	var count int64
	err := r.db.WithContext(ctx).Where("user_name = ?", userName).Count(&count).Error
	if err != nil {
		fmt.Println(err)
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r *userRepository) DeleteUser(ctx context.Context, userName string) error {
	return r.db.WithContext(ctx).Where("user_name = ?", userName).Delete(&user.User{}).Error
}
