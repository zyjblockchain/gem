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
		serializer.ErrorResponse(c, 40001, "需要先登录", "")
		// 终止后面的handle执行
		c.Abort()
	}
}
