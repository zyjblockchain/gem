package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, exist := c.Get("user"); exist {
			if _, ok := user.(*models.User); ok {
				c.Next()
				return
			}
		}

		// 需要先登录
		c.JSON(200, serializer.Response{
			Status: 40001,
			Data:   nil,
			Msg:    "需要先登录",
			Error:  "",
		})
		// 终止后面的handle执行
		c.Abort()
	}
}
