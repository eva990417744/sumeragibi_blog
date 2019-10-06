package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sumeragibi_blog/log_init"
	"sumeragibi_blog/modes"
	"sumeragibi_blog/redis_cline"
)

var log *zap.Logger

var redisCline *redis.Client
var dbCline *gorm.DB

func ProjectInit() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		log.Error(err.Error())
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	redisAddr := viper.GetString("redis.addr")
	redisPassword := viper.GetString("redis.password")
	redisDb := viper.GetInt("redis.db")
	sqlHost := viper.GetString("postgres.host")
	sqlPort := viper.GetString("postgres.port")
	sqlUser := viper.GetString("postgres.user")
	sqlDb := viper.GetString("postgres.dbname")
	sqlPass := viper.GetString("postgres.password")
	log = log_init.LogInit()
	redisCline = redis_cline.NewRedisClient(redisAddr, redisPassword, redisDb)
	dbCline = modes.DataBaseCline(sqlHost, sqlPort, sqlUser, sqlDb, sqlPass)
}

func main() {
	ProjectInit()
	redisCline.Ping()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
