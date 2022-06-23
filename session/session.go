package session

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	otpSessionTTLSecond = 5 * 60
	sessionTTLSecond    = 60 * 60
)

type SessionClient struct {
	rdb *redis.Client
}

func NewClient() *SessionClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &SessionClient{rdb: rdb}
}

func (sc SessionClient) NewSessionId() (string, error) {
	length := 36
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (sc SessionClient) CreateOTPSession(ctx context.Context, userId string) (string, error) {
	sid, err := sc.NewSessionId()
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("sessions:otp:%s", sid)
	if err := sc.Set(ctx, s, userId, otpSessionTTLSecond*time.Second); err != nil {
		return "", err
	}
	return s, nil
}

func (sc SessionClient) GetOTPSession(ctx context.Context, sessionId string) (userId string, err error) {
	userId, err = sc.Get(ctx, fmt.Sprintf("session:otp:%s", sessionId))
	if err != nil {
		return "", err
	}
	return userId, err
}

func (sc SessionClient) CreateLoginSession(ctx context.Context, userId string) (string, error) {
	sid, err := sc.NewSessionId()
	if err != nil {
		return "", err
	}

	oKey := fmt.Sprintf("sessions:otp:%s", sid)
	if err := sc.Del(ctx, oKey); err != nil {
		return "", err
	}

	skey := fmt.Sprintf("sessions:login:%s", sid)
	if err := sc.Set(ctx, skey, userId, sessionTTLSecond*time.Second); err != nil {
		return "", err
	}
	return skey, nil
}

func (sc SessionClient) GetLoginSession(ctx context.Context, sessionId string) (userId string, err error) {
	userId, err = sc.Get(ctx, fmt.Sprintf("session:login:%s", sessionId))
	if err != nil {
		return "", err
	}
	return userId, err
}

func (sc SessionClient) Get(ctx context.Context, key string) (string, error) {
	return sc.rdb.Get(ctx, key).Result()
}

func (sc SessionClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return sc.rdb.Set(ctx, key, value, expiration).Err()
}

func (sc SessionClient) Del(ctx context.Context, key string) error {
	return sc.rdb.Del(ctx, key).Err()
}
