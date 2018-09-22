package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
)

type UrlMap struct {
	Id 				int64 `grom:"primary_key AUTO_INCREMENT"`
	Url 			string `grom:"size:9000"`
	CreateTime 		int64 `grom:"column:create_time"`
}

// 表名设置
func (UrlMap) TableName() string {
	return "urlmap"
}

func GenMURL(user, passwd, host, port, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, host, port, database)
}

func AddUrl(db *gorm.DB, url string) int64 {
	id := IdGenerator(1)
	urlMap := UrlMap{id, url, time.Now().Unix()}

	db.Create(&urlMap)
	return id
}

func SelectUrl(db *gorm.DB, id int64) string {
	var res UrlMap
	db.Where("id = ?", id).Find(&res)

	if res.Id == 0 {
		return ""
	}

	return res.Url
}
