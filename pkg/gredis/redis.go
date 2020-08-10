package gredis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yrjkqq/tiny-website/pkg/setting"
)

// RDB redis client
var RDB *redis.Client

var ctx = context.Background()

// Setup set up redis
func Setup() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:        setting.RedisSetting.Addr,
		Password:    setting.RedisSetting.Password,
		Network:     setting.RedisSetting.Network,
		DialTimeout: setting.RedisSetting.DialTimeout * time.Second,
		DB:          0,
		OnConnect: func(ctx context.Context, c *redis.Conn) error {
			pong, err := c.Ping(ctx).Result()
			log.Println(pong)
			return err
		},
	})
	r, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Redis conn test failed, result: %v, error: %v", r, err)
	}
	return nil
}

// Set ...
func Set(key string, data interface{}, t time.Duration) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = RDB.Set(ctx, key, value, t).Err()
	return err
}

// Exists ...
func Exists(key string) bool {
	return RDB.Exists(ctx, key).Val() == 1
}

// Get ...
func Get(key string) (string, error) {
	val, err := RDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Del ...
func Del(key string) (bool, error) {
	val, err := RDB.Del(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}

// LikeDels ...
func LikeDels(key string) error {
	keys, err := RDB.Keys(ctx, fmt.Sprintf("*%v*", key)).Result()
	if err != nil {
		return err
	}
	_, err = RDB.Del(ctx, keys...).Result()
	return err
}
