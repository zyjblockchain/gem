package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/zyjblockchain/gem/serializer"
	"net/http"
)

// Ping 用于心跳检测
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Status: http.StatusOK,
		Data:   nil,
		Msg:    "Pong",
		Error:  "",
	})
}
