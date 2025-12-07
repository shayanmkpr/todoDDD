package auth

import "os"

var (
	AccessTokenSecret  string = os.Getenv("ACCESS_TOKEN_SECRET")
	RefreshTokenSecret string = os.Getenv("REFRESH_TOKEN_SECRET")
)
