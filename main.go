package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/Unknwon/goconfig"
)

var globalDB *gorm.DB

func init()  {
	// 加载配置
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		fmt.Printf("config parse error, %s", err.Error())
	}
	mysqlConf, err := cfg.GetSection("mysql")
	if err != nil {
		fmt.Printf("config parse error, %s", err.Error())
	}
	// 数据库连接
	dbDSN := GenMURL(mysqlConf["user"], mysqlConf["passwd"], mysqlConf["host"], mysqlConf["port"], mysqlConf["database"])
	globalDB, err = gorm.Open("mysql", dbDSN)
	if err != nil {
		fmt.Printf("error in db connect, %s", err.Error())
	}
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
