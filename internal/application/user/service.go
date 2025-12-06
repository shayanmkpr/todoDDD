package user

import (
	"context"
	"errors"

	"todoDB/internal/domain/user"
)

type Service struct { // initiated a repo that is already defined.
	repo user.UserRepository
}

func NewService(r user.UserRepository) *Service {
	return &Service{repo: r}
}

// Register registers a new user.
func (s *Service) Register(ctx context.Context, userName, pass string) (*user.User, error) {
	// check if user exists
	userNameExists, err := s.repo.CheckByName(ctx, userName)
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
	if err := s.repo.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) Login(ctx context.Context, userName, pass string) (string, error) {
	token := "token"
	// Get the user
	user, err := s.repo.GetByNme(ctx, userName)
	if err != nil {
		return "", err
	}

	correctPass, err := user.CheckPassword(pass)
	if err != nil {
		return "", err
	}

	if correctPass {
		// generate the token
		return token, nil
	}
	return token, nil
}
