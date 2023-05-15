package config

import (
	"Low-Battery-Inquiry-Send/config/redis"
	"context"
	"time"
)

var ctx = context.Background()

func GetAccessTokenKey() (string, error) {
	val, err := redis.RedisClient.Get(ctx, "accessTokenKey").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func SetAccessTokenKey(value string, ExpiresIn int) {
	redis.RedisClient.Set(ctx, "accessTokenKey", value, time.Duration(ExpiresIn)*time.Second)
}

func CheckAccessTokenKey() bool {
	intCmd := redis.RedisClient.Exists(ctx, "accessTokenKey")
	if intCmd.Val() == 1 {
		return true
	} else {
		return false
	}
}
