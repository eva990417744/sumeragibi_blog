package main

import (
	"sumeragibi_blog/log_init"
	"sumeragibi_blog/redis_cline"
)

var redisCline = redis_cline.NewRedisClient()
var log  = log_init.LogInit()


func main() {
	redisCline.Ping()
	log.Error("hhh")
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//_ = r.Run() // listen and serve on 0.0.0.0:8080
}
