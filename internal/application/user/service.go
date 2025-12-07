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

// Register registers a new user. But we are not giving them a token right away.
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
	if err := s.repo.SaveUser(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) Login(ctx context.Context, userName, pass string) (string, error) {
	// Get the user
	user, err := s.repo.GetByName(ctx, userName)
	if err != nil {
		return "", err
	}

	isCorrectPass, err := user.CheckPassword(pass)
	if err != nil {
		return "", err
	}

	if isCorrectPass {
		// generate the token
		return token, nil
	}
	return token, nil
}
