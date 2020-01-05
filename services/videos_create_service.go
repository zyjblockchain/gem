package services

import (
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
)

type CreateVideosInfo struct {
	Title      string `form:"title" json:"title" binding:"required,min=3,max=1000"` // 视频标题
	Info       string `form:"info" json:"info" binding:"max=30000"`                 // 视频简介
	VideoPath  string `form:"videoPath" json:"videoPath" binding:"required"`        // oss视频源url
	AvatarPath string `form:"avatarPath" json:"avatarPath" binding:"required"`      // oss视频封面url
}

// CreateVideos 保存前端传入的视频信息到数据库并返回oss授权的关于视频put 到oss的url，前端通过此url直传视频到oss进行存储
func (c *CreateVideosInfo) CreateVideos(userId string) (*serializer.VideoPut, error) {
	video := &models.Video{
		Title:      c.Title,
		Info:       c.Info,
		VideoPath:  c.VideoPath,
		AvatarPath: c.AvatarPath,
		OwnerId:    userId,
	}

	// 保存数据库
	err, newVideo := models.AddVideo(video)
	if err != nil {
		return nil, err
	}

	// 组装返回的结果
	videoPuts, err := serializer.AssembleVideoPut([]*models.Video{newVideo})
	if err != nil || len(videoPuts) == 0 {
		return nil, err
	}

	return videoPuts[0], nil
}
