package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
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
	ttl := time.Until(token.ExpiresAt)
	key := "refresh:" + token.Value
	// lua script for redis atomicity. This is cool.
	script := redis.NewScript(`
		redis.call("HSET", KEYS[1],
			"value", ARGV[1],
			"userName", ARGV[2],
			"expiresAt", ARGV[3])
		redis.call("EXPIRE", KEYS[1], ARGV[4])
		return 1
	`)
	ttlSeconds := int(ttl.Seconds())
	_, err := script.Run(ctx, t.rdb, []string{key}, token.Value, token.UserName, token.ExpiresAt.Unix(), ttlSeconds).Result()
	if err != nil {
		return err
	}
	return nil
}

func (t *refreshTokenRepo) GetRefreshToken(ctx context.Context, tokenValue string) (*auth.RefreshToken, error) {
	key := "refresh:" + tokenValue
	hashToken, err := t.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if len(hashToken) == 0 { // what?
		fmt.Println(hashToken)
		return nil, errors.New("Record not set correctly")
	}

	//convert time to time.Time
	expUnix, err := strconv.ParseInt(hashToken["expiresAt"], 10, 64)

	return &auth.RefreshToken{
		Value:     hashToken["value"],
		UserName:  hashToken["userName"],
		ExpiresAt: time.Unix(expUnix, 0),
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
