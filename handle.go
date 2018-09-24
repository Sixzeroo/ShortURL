package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	if !isUrl(url) {
		c.JSON(200, gin.H{
			"code": 1,
			"message": "Invaild url",
		})
		return
	}

	id  := AddUrl(globalDB, url)
	idStr := Int2Str(id)

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

	id := Str2Int(idStr)
	url := SelectUrl(globalDB, id)

	// 使用302重定向
	c.Redirect(http.StatusFound, url)
}
