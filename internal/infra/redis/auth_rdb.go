package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"todoDB/internal/domain/auth"

	"github.com/redis/go-redis/v9"
)

type refreshTokenRepo struct {
	rdb *redis.Client
}

func NewRdbRepository(rdb *redis.Client) auth.RefreshTokenRepo {
	return &refreshTokenRepo{rdb: rdb}
}

func (t *refreshTokenRepo) StoreRefreshToken(ctx context.Context, token *auth.RefreshToken) error {
	fmt.Println("storing stuff in redis")
	defer fmt.Println("done storing stuff in redis")

	ttl := time.Until(token.ExpiresAt)
	ttlSeconds := int(ttl.Seconds())

	tokenKey := "refresh:" + token.Value
	userKey := "user:" + token.UserName + ":refresh"

	// this should be checked. There should be two main relations, one with user object and a key, and another with the key and the user name of a registered user object.
	script := redis.NewScript(`
		redis.call("HSET", KEYS[1],
			"value", ARGV[1],
			"userName", ARGV[2],
			"expiresAt", ARGV[3]
		)
		-- storing name as the key and the token as the value
		redis.call("HSET", KEYS[2],
			"token", ARGV[1]
		)
		redis.call("EXPIRE", KEYS[1], ARGV[4])
		redis.call("EXPIRE", KEYS[2], ARGV[4])
		return 1
	`)

	_, err := script.Run(
		ctx,
		t.rdb,
		[]string{tokenKey, userKey}, // KEYS[1], KEYS[2]
		token.Value,                 // ARGV[1]
		token.UserName,              // ARGV[2]
		token.ExpiresAt.Unix(),      // ARGV[3]
		ttlSeconds,                  // ARGV[4]
	).Result()

	return err
}

func (t *refreshTokenRepo) GetRefreshToken(ctx context.Context, tokenValue string) (string, error) {
	key := "refresh:" + tokenValue
	hashToken, err := t.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return "", err
	}

	if len(hashToken) == 0 { // what?
		return "", errors.New("Record not set correctly")
	}

	return hashToken["value"], nil
}

// make sure the token coming from redis is actually a string
func (t *refreshTokenRepo) GetRefreshTokenByName(ctx context.Context, userName string) (*auth.RefreshToken, error) {
	key := "user:" + userName
	token, err := t.rdb.Get(ctx, key).Result() // token is a string
	if err != nil {
		return nil, err
	}

	if len(token) == 0 { // what?
		return nil, errors.New("Record not set correctly")
	}

	return &auth.RefreshToken{
		Value:     token,
		UserName:  userName,
		ExpiresAt: time.Now(),
	}, nil
}

func (t *refreshTokenRepo) DeleteRefershToken(ctx context.Context, tokenValue string) error {
	key := "refresh:" + tokenValue
	_, err := t.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}
	t.rdb.Del(ctx, key)
	return nil
}
