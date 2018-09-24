package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func IndexHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func AddHandle(c *gin.Context) {
	url := c.DefaultPostForm("url", "")

	if url == "" {
		c.JSON(200, gin.H{
			"code": 1,
			"message": "Invaild url",
		})
		return
	}

	// url 合法性验证
	if !isUrl(url) {
		c.JSON(200, gin.H{
			"code": 1,
			"message": "Invaild url",
		})
		return
	}

	// long url -> short url 缓存
	idStrValue, err := redisClient2.Get(url).Result()
	if err != nil {
		fmt.Printf("error in get redis value : %s\n", err.Error())
	}
	if idStrValue != "" {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "Success",
			"id_str": idStrValue,
		})
		return
	}

	id  := AddUrl(globalDB, url)
	idStr := Int2Str(id)

	err = redisClient2.Set(url, idStr, time.Duration(3600 * 24 * 10) * time.Second).Err()
	if err != nil {
		fmt.Printf("error in set redis value : %s\n", err.Error())
	}

	c.JSON(200, gin.H{
		"code": 0,
		"message": "Success",
		"id_str": idStr,
	})
}

func ParseHandle(c *gin.Context) {
	idStr := c.Param("idstr")

	if !isIdStr(idStr) {
		c.JSON(200, gin.H{
			"code": 1,
			"message": "Invaild url",
		})
		return
	}

	// short url -> long url 缓存
	urlValue, err := redisClient1.Get(idStr).Result()
	if err != nil {
		fmt.Printf("error in get redis value : %s\n", err.Error())
	}
	if urlValue != "" {
		// 使用302重定向
		c.Redirect(http.StatusFound, urlValue)
		return
	}

	id := Str2Int(idStr)
	url := SelectUrl(globalDB, id)

	// 过期时间设置为10天
	err = redisClient1.Set(idStr, url, time.Duration(3600 * 24 * 10) * time.Second).Err()
	if err != nil {
		fmt.Printf("error in set redis value : %s\n", err.Error())
	}

	// 使用302重定向
	c.Redirect(http.StatusFound, url)
}
