package redis_cline

import (
	"github.com/go-redis/redis"
)



func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"", // no password set
		DB:0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err !=nil{
		println(pong)
		return nil
	}
	return client
}

