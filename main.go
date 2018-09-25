package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/go-redis/redis"
	"log"
)

var globalDB *gorm.DB
var redisClient1, redisClient2 *redis.Client


func init()  {
	// 加载配置
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		log.Printf("config parse error, %s", err.Error())
	}
	mysqlConf, err := cfg.GetSection("mysql")
	if err != nil {
		log.Printf("config parse error, %s", err.Error())
	}
	// 数据库连接
	dbDSN := GenMURL(mysqlConf["user"], mysqlConf["passwd"], mysqlConf["host"], mysqlConf["port"], mysqlConf["database"])
	globalDB, err = gorm.Open("mysql", dbDSN)
	if err != nil {
		log.Printf("error in db connect, %s", err.Error())
	}
	// Redis连接
	redisConf, err := cfg.GetSection("redis")
	if err != nil {
		log.Printf("redis config parse error %s", err.Error())
	}
	redisClient1 = redis.NewClient(&redis.Options{
		Addr:	fmt.Sprintf("%s:%s", redisConf["host"], redisConf["port"]),
		Password: "",
		DB: 0,
	})
	redisClient2 = redis.NewClient(&redis.Options{
		Addr:	fmt.Sprintf("%s:%s", redisConf["host"], redisConf["port"]),
		Password: "",
		DB: 1,
	})
}

func main() {
	r := gin.Default()

	//r.GET("/ping", Ping)

	r.GET("/:idstr", ParseHandle)
	r.POST("/", AddHandle)
	r.GET("/", IndexHandle)

	r.LoadHTMLGlob("templates/*")
	r.Run()
}
