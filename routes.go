package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

const patternReg = `^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`

// 建立連結
func createUrl(c *gin.Context) {
	var b struct {
		Url string `json:"url" binding:"required"`
	}

	err := c.ShouldBindJSON(&b)
	if err != nil {
		// 結構有問題
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
		return
	}

	ok, _ := regexp.MatchString(patternReg, b.Url)
	if !ok {
		// 不是網址
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
		})
		return
	}

	path := Uuid()
	Db.Set(path, b.Url)

	c.JSON(http.StatusOK, gin.H{
		"url":  path,
		"code": 0,
	})

	log.Printf("- %s - %s\n", c.ClientIP(), b.Url)
}

// 主頁
func index(c *gin.Context) {
	b, _ := Fs.ReadFile("public/index.html")
	c.Writer.Write(b)
}

// 找不到
func noRoute(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

// 重新導向
func redirect(c *gin.Context) {
	key := c.Param("key")
	url, ok := Db.Get(key)
	if ok {
		c.Redirect(http.StatusFound, url)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}
