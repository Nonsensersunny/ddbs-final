package redis

import (
	"ddbs-final/internal/config"
	"fmt"
	"github.com/go-redis/redis"
)

func GetRedisClient(client *config.RedisConf) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", client.Host, client.Port),
		Password: client.Password,
		DB:       client.DBName,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		panic(err.Error())
	}
	return redisClient
}
