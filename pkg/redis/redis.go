package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type IRedis interface {
	StoreKeyValue(key string, value string, expiration time.Duration)
	GetKeyValue(key string) (string, error)
	DeleteKey(key string) *redis.IntCmd
}

type ObjectStorageWrapper struct {
	InitialiseRedisClient *redis.Client
}

func (dw *ObjectStorageWrapper) StoreKeyValue(key string, value string, expiration time.Duration) {
	err := dw.InitialiseRedisClient.Set(key, value, expiration).Err()
	if err != nil {
		panic(errors.New("Could not set key-value pair in Redis. Please check your connection"))
	}
}

func (dw *ObjectStorageWrapper) GetKeyValue(key string) (string, error) {
	val, err := dw.InitialiseRedisClient.Get(key).Result()
	return val, err
}

func (dw *ObjectStorageWrapper) DeleteKey(key string) *redis.IntCmd {
	done := dw.InitialiseRedisClient.Del(key)
	return done
}

func (dw *ObjectStorageWrapper) ExistsKey(key string) *redis.IntCmd {
	exists := dw.InitialiseRedisClient.Exists(key)
	return exists
}
