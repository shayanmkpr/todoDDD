package user

import (
	"errors"
	"log"
)

type User struct {
	UserName string `json:"name" binding:"required"`
	Pass     string `json:"pass" binding:"required"`
}

func (u *User) UpdateUserName(newUserName string) error {
	if u.UserName == "" {
		log.Println("the UserName is empty")
		return errors.New("the UserName is empty")
	}
	u.UserName = newUserName
	return nil
}
