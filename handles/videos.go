package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
	"github.com/zyjblockchain/gem/services"
	"strconv"
)

// 获取自己上传的视频
func GetOwnerVideos() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.OwnerVideos
		if err := c.ShouldBind(&service); err == nil {
			u, _ := c.Get("user")
			videos, err := service.GetVideos(u.(*models.User).ID)
			if err != nil {
				serializer.ErrorResponse(c, 40001, "获取自己的视频失败", err.Error())
			} else {
				serializer.SuccessResponse(c, videos, "success")
			}
		} else {

			serializer.ErrorResponse(c, 40001, "参数错误", err.Error())

		}
	}
}

// CreateVideo 上传视频
func CreateVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := services.CreateVideosInfo{}
		if err := c.ShouldBind(&service); err == nil {
			// 保存视频
			// 获取当前用户
			u, exist := c.Get("user")
			if !exist {
				serializer.ErrorResponse(c, 40001, "请先登录", "")
			}
			user, ok := u.(*models.User)
			if !ok {
				panic("user保存的登录态不正确")
			}

			video, err := service.CreateVideos(strconv.Itoa(int(user.ID)))
			if err != nil {
				serializer.ErrorResponse(c, 40001, "创建视频失败", err.Error())
			} else {
				serializer.SuccessResponse(c, *video, "视频创建成功")
			}
		} else {
			serializer.ErrorResponse(c, 40001, "参数错误", err.Error())

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
				serializer.ErrorResponse(c, 40001, "show视频失败", err.Error())
			} else {
				serializer.SuccessResponse(c, *video, "show视频成功")
			}
		} else {
			serializer.ErrorResponse(c, 40001, "参数错误", err.Error())
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
				serializer.ErrorResponse(c, 40001, "ListVideo 失败", err.Error())
			} else {
				serializer.SuccessResponse(c, videos, "成功")
			}
		} else {
			serializer.ErrorResponse(c, 40001, "参数错误", err.Error())
		}
	}
}

func UpdateVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.UpdateVideoInfo
		if err := c.ShouldBind(&service); err == nil {
			newVideo, err := service.UpdateVideo(c.Param("id"))
			if err != nil {
				serializer.ErrorResponse(c, 40001, "更新视频失败", err.Error())
			} else {
				serializer.SuccessResponse(c, *newVideo, "成功")
			}
		} else {
			serializer.ErrorResponse(c, 40001, "参数错误", err.Error())
		}
	}
}

func DeleteVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.DeleteVideoInfo
		if err := c.ShouldBind(&service); err == nil {
			if err := service.DelVideo(c.Param("id")); err != nil {
				serializer.ErrorResponse(c, 40001, "删除视频失败", err.Error())

			} else {
				serializer.SuccessResponse(c, nil, "成功")
			}
		} else {
			serializer.ErrorResponse(c, 40001, "参数错误", err.Error())
		}
	}
}
