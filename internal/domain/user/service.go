package user

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) UpdateUserName(newUserName string) error {
	if u.UserName == "" {
		log.Println("the UserName is empty")
		return errors.New("the UserName is empty")
	}
	u.UserName = newUserName
	return nil
}

func NewUser(userName, pass string) (*User, error) {
	if userName == "" || pass == "" {
		return nil, errors.New("the username or password can't be empty")
	}
	newUser := &User{
		UserName: userName,
		Pass:     pass,
	}

	return newUser, nil
}

func (u *User) CheckPassword(inputPass string) (bool, error) {
	if inputPass == "" {
		return false, errors.New("input pass is empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Pass), []byte(inputPass))
	if err != nil {
		fmt.Println(err)
		// password mismatch (not a server error)
		return false, nil
	}

	return true, nil
}
