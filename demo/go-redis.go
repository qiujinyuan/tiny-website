package demo

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/yrjkqq/tiny-website/pkg/setting"
)

var ctx = context.Background()

// GoRedisExampleClient go-redis example
func GoRedisExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:        setting.RedisCfg["addr"].(string),
		Password:    setting.RedisCfg["password"].(string),
		Network:     setting.RedisCfg["network"].(string),
		DialTimeout: time.Duration(setting.RedisCfg["dialtimeout"].(int)) * time.Second,
		DB:          0,
		OnConnect: func(ctx context.Context, c *redis.Conn) error {
			pong, err := c.Ping(ctx).Result()
			fmt.Println(pong)
			return err
		},
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// return 1 if exists
	fmt.Println(rdb.Exists(ctx, "key").Val())

	fmt.Println(rdb.Del(ctx, "key2").Val())

	fmt.Println(rdb.Exists(ctx, "key").Val())

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	rdb.Set(ctx, "key1", "valuekey1", 0)
	rdb.Set(ctx, "no", "valueno", 0)
	rdb.Set(ctx, "key2", "valuekey2", 0)
	rdb.Set(ctx, "key3", "valuekey3", 0)
	rdb.Set(ctx, "2key4", "value2key4", 0)
	keys, _ := rdb.Keys(ctx, "*").Result()
	// for _, key := range keys {
	// 	fmt.Println(rdb.Get(ctx, key).Result())
	// }

	fmt.Println(rdb.Del(ctx, keys...).Result())
}
