package application

import (
	"context"
	"errors"
	"os"
	"time"

	"todoDB/internal/domain/auth"
	"todoDB/internal/domain/user"
)

var (
	AccessTokenSecret  string = os.Getenv("ACCESS_TOKEN_SECRET")
	RefreshTokenSecret string = os.Getenv("REFRESH_TOKEN_SECRET")
)

type UserService struct { // initiated a repo that is already defined.
	refreshRepo auth.RefreshTokenRepo
	authRepo    auth.AuthenticationRepo
	userRepo    user.UserRepository
}

func NewService(r user.UserRepository, a auth.AuthenticationRepo, ref auth.RefreshTokenRepo) *UserService {
	return &UserService{
		refreshRepo: ref,
		userRepo:    r,
		authRepo:    a,
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
		return nil, errors.New("user name already registered")
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

func (s *UserService) Login(ctx context.Context, userName, pass string) (string, string, error) {
	// Get the user
	user, err := s.userRepo.GetByName(ctx, userName)
	if err != nil {
		return "", "", err
	}

	isCorrectPass, err := user.CheckPassword(pass)
	if err != nil {
		return "", "", err
	}

	if isCorrectPass {
		accessToken, err := s.authRepo.GenerateAccessToken(ctx, AccessTokenSecret, userName)
		if err != nil {
			return "", "", err
		}

		refreshTken, err := s.authRepo.GenerateRefreshToken(ctx, RefreshTokenSecret, userName)
		if err != nil {
			return "", "", err
		}

		// here should go the stroeRefreshToken logic

		return accessToken, refreshTken, nil
	} else {
		return "", "", errors.New("the password is incorrect")
	}
}

// check if the token is correctly signed?
func (s *UserService) TokenLogin(ctx context.Context, refreshToken string) (string, error) {
	inputCalimsPtr, err := s.authRepo.ParseToken(ctx, RefreshTokenSecret, refreshToken)
	if err != nil {
		return "", err
	}
	// check with redis. --> call the auth_redis_repo
	refreshTokenParsed, err := s.refreshRepo.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	if refreshTokenParsed.ExpiresAt.Before(time.Now()) {
		return "", errors.New("the refresh token has expired. Need to login manually again")
	}
	userName := inputCalimsPtr.UserName
	newToken, err := s.authRepo.GenerateAccessToken(ctx, AccessTokenSecret, userName)
	if err != nil {
		return "", err
	}
	return newToken, nil
}
