package redis_cline

import (
	"fmt"
	"github.com/go-redis/redis"
	"sumeragibi_blog/log_init"
)

var log  = log_init.LogInit()


func NewRedisClient(addr string,password string,db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"", // no password set
		DB:0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err !=nil{
		log.Error(err.Error())
		panic(fmt.Errorf("Ping redis error: %s \n", err))
	}
	fmt.Printf(pong)
	return client
}

