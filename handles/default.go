package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/zyjblockchain/gem/serializer"
)

// Ping 用于心跳检测
func Ping(c *gin.Context) {
	serializer.SuccessResponse(c, nil, "pong")

}
