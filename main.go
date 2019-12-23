package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化一个http服务
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8001")
}
