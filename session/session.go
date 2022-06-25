package session

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"math/big"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	otpSessionTTLSecond = 5 * 60
	sessionTTLSecond    = 60 * 60
	// store used otp only 2 codes (1 code: 30 sec)
	usedOtpSecond = 60
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

func (sc SessionClient) NewSessionId() string {
	length := 36
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(255))
		b[i] = byte(num.Int64())
	}

	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	return encoder.EncodeToString(b)
}

func (sc SessionClient) CreateOTPSession(ctx context.Context, userId string) (string, error) {
	sid := sc.NewSessionId()

	s := fmt.Sprintf("sessions:otp:%v", sid)
	if err := sc.Set(ctx, s, userId, otpSessionTTLSecond*time.Second); err != nil {
		return "", err
	}
	return s, nil
}

func (sc SessionClient) GetOTPSession(ctx context.Context, sessionId string) (userId string, err error) {
	userId, err = sc.Get(ctx, sessionId)
	if err != nil {
		return "", err
	}
	return userId, err
}

func (sc SessionClient) CreateLoginSession(ctx context.Context, userId string) (string, error) {
	sid := sc.NewSessionId()
	oKey := fmt.Sprintf("sessions:otp:%v", sid)
	if err := sc.Del(ctx, oKey); err != nil {
		return "", err
	}

	skey := fmt.Sprintf("sessions:login:%v", sid)
	if err := sc.Set(ctx, skey, userId, sessionTTLSecond*time.Second); err != nil {
		return "", err
	}
	return skey, nil
}

func (sc SessionClient) GetLoginSession(ctx context.Context, sessionId string) (userId string, err error) {
	userId, err = sc.Get(ctx, sessionId)
	if err != nil {
		return "", err
	}
	return userId, err
}

func (sc SessionClient) SetUsedOTP(ctx context.Context, userId string, otp string) error {
	key := fmt.Sprintf("otp:used:%v", userId)
	if err := sc.LPush(ctx, key, otp); err != nil {
		return err
	}

	// will store 2 otps
	if err := sc.LTrim(ctx, key, 0, 1); err != nil {
		return err
	}

	return sc.Expire(ctx, key, usedOtpSecond*time.Second)
}

func (sc SessionClient) GetUsedOTPs(ctx context.Context, userId string) []string {
	key := fmt.Sprintf("otp:used:%v", userId)
	return sc.rdb.LRange(ctx, key, 0, 1).Val()
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

func (sc SessionClient) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return sc.rdb.Expire(ctx, key, ttl).Err()
}

func (sc SessionClient) LPush(ctx context.Context, key string, values ...interface{}) error {
	return sc.rdb.LPush(ctx, key, values...).Err()
}

func (sc SessionClient) LTrim(ctx context.Context, key string, start int64, stop int64) error {
	return sc.rdb.LTrim(ctx, key, start, stop).Err()
}

func (sc SessionClient) LPop(ctx context.Context, key string, values ...interface{}) (string, error) {
	return sc.rdb.LPop(ctx, key).Result()
}
