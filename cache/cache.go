package cache

import (
	"context"
	"fmt"
	"go-image/config"
	"go-image/convert"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client
var maxCache uint
var expireTime uint
var IsCache bool
var ctx = context.Background()

func init() {

	IsCache = convert.StringToBool(config.GetSetting("redis.cache"))
	maxCache = convert.StringToUint(config.GetSetting("redis.max_cache"))
	expireTime = convert.StringToUint(config.GetSetting("redis.expire_time"))

	if IsCache {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.GetSetting("redis.addr"),
			Password: config.GetSetting("redis.password"),
			DB:       convert.StringToInt(config.GetSetting("redis.db")),
		})

		if _, err := redisClient.Ping(ctx).Result(); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

}

func Set(key string, value interface{}) {
	if uint(len(value.([]byte))) <= maxCache {
		err := redisClient.Set(ctx, key, value, time.Second*time.Duration(expireTime)).Err()
		if err != nil {
			log.Println(err)
		}
	}

}

func Get(key string) *[]byte {
	val, err := redisClient.Get(ctx, key).Bytes()
	if err == redis.Nil || err == nil {
		return &val
	}

	return nil
}

func Del(key string) {
	vals, err := redisClient.Keys(ctx, key+":*").Result()
	if err == redis.Nil || err == nil {
		redisClient.Del(ctx, vals...)
	}
}
