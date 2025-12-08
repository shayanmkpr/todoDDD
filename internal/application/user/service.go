package application

import (
	"context"
	"errors"
	"os"

	"todoDB/internal/domain/auth"
	"todoDB/internal/domain/user"
)

var (
	AccessTokenSecret  string = os.Getenv("ACCESS_TOKEN_SECRET")
	RefreshTokenSecret string = os.Getenv("REFRESH_TOKEN_SECRET")
)

type UserService struct { // initiated a repo that is already defined.
	userRepo user.UserRepository
	authRepo auth.AuthenticationRepo
}

func NewService(r user.UserRepository, a auth.AuthenticationRepo) *UserService {
	return &UserService{
		userRepo: r,
		authRepo: a,
	}
}

// Register registers a new user. But we are not giving them a token right away.
func (s *UserService) Register(ctx context.Context, userName, pass string) (*user.User, error) {
	// check if user exists
	userNameExists, err := s.userRepo.CheckByName(ctx, userName)
	if err != nil {
		return nil, err
	}
	if userNameExists {
		return nil, errors.New("email already registered")
	}

	// domain creates the user object
	u, err := user.NewUser(userName, pass) // lacks the id yet.
	if err != nil {
		return nil, err
	}

	// store it and give it an ID
	if err := s.userRepo.SaveUser(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) Login(ctx context.Context, userName, pass string) (string, error) {
	// Get the user
	user, err := s.userRepo.GetByName(ctx, userName)
	if err != nil {
		return "", err
	}

	isCorrectPass, err := user.CheckPassword(pass)
	if err != nil {
		return "", err
	}

	if isCorrectPass {
		token, err := s.authRepo.GenerateAccessToken(ctx, AccessTokenSecret, userName)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", errors.New("The password is incorrect")
	}
}
