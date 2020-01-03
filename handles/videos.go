package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/zyjblockchain/gem/serializer"
	"github.com/zyjblockchain/gem/services"
)

// 获取自己上传的视频
func GetOwnerVideos() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.OwnerVideos
		if err := c.ShouldBind(&service); err == nil {
			videos, err := service.GetVideos(c.Param("user_id"))
			if err != nil {
				c.JSON(200, serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "获取owner videos error",
					Error:  err.Error(),
				})
			} else {
				c.JSON(200, serializer.Response{
					Status: 0,
					Data:   videos,
					Msg:    "success",
					Error:  "",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "参数错误",
				Error:  err.Error(),
			})
		}
	}
}

// CreateVideo 上传视频
func CreateVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := services.CreateVideosInfo{}
		if err := c.ShouldBind(&service); err == nil {
			// 保存视频
			video, err := service.CreateVideos(c.Param("user_id"))
			if err != nil {
				c.JSON(200, serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "创建视频失败",
					Error:  err.Error(),
				})
			} else {
				c.JSON(200, serializer.Response{
					Status: 200,
					Data:   *video,
					Msg:    "视频创建成功",
					Error:  "",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "参数错误",
				Error:  err.Error(),
			})
		}
	}
}

// ShowVideo
func ShowVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.ShowVideoInfo
		if err := c.ShouldBind(&service); err == nil {
			video, err := service.ShowVideo(c.Param("id"))
			if err != nil {
				c.JSON(200, serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "show视频失败",
					Error:  err.Error(),
				})
			} else {
				c.JSON(200, serializer.Response{
					Status: 200,
					Data:   *video,
					Msg:    "show视频成功",
					Error:  "",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "参数错误",
				Error:  err.Error(),
			})
		}
	}
}

// 获取当前账户的所有视频
func ListVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.BatchVideosInfo
		if err := c.ShouldBind(&service); err == nil {
			videos, err := service.BatchVideos()
			if err != nil {
				c.JSON(200, serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "ListVideo 失败",
					Error:  err.Error(),
				})
			} else {
				c.JSON(200, serializer.Response{
					Status: 200,
					Data:   videos,
					Msg:    "成功",
					Error:  "",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "参数错误",
				Error:  err.Error(),
			})
		}
	}
}

func UpdateVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.UpdateVideoInfo
		if err := c.ShouldBind(&service); err == nil {
			newVideo, err := service.UpdateVideo(c.Param("id"))
			if err != nil {
				c.JSON(200, serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "更新视频失败",
					Error:  err.Error(),
				})
			} else {
				c.JSON(200, serializer.Response{
					Status: 200,
					Data:   *newVideo,
					Msg:    "成功",
					Error:  "",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "参数错误",
				Error:  err.Error(),
			})
		}
	}
}

func DeleteVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.DeleteVideoInfo
		if err := c.ShouldBind(&service); err == nil {
			if err := service.DelVideo(c.Param("id")); err != nil {
				c.JSON(200, serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "删除视频失败",
					Error:  err.Error(),
				})
			} else {
				c.JSON(200, serializer.Response{
					Status: 200,
					Data:   nil,
					Msg:    "成功",
					Error:  "",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "参数错误",
				Error:  err.Error(),
			})

		}
	}
}
