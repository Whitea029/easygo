package cache

import (
	"{{ .GoModule }}/conf"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func InitCache() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%d", conf.GetConfig().Redis.Host, conf.GetConfig().Redis.Port),
		DB:                 conf.GetConfig().Redis.DB,
		DialTimeout:        conf.GetConfig().Redis.DialTimeout,
		ReadTimeout:        conf.GetConfig().Redis.ReadTimeout,
		WriteTimeout:       conf.GetConfig().Redis.WriteTimeout,
		PoolSize:           conf.GetConfig().Redis.PoolSize,
		PoolTimeout:        conf.GetConfig().Redis.PoolTimeout,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: conf.GetConfig().Redis.IdleCheckFrequency * time.Millisecond,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return
}

func CloseCache() {
	redisClient.Close()
}

func Set[T any](c *redis.Client, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(key, v, duration).Err()
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var dest T = *new(T)
	v, err := c.Get(key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}
