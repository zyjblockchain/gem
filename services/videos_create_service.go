package services

import (
	"github.com/zyjblockchain/gem/models"
)

type CreateVideosInfo struct {
	Title string `form:"title" json:"title" binding:"required,min=3,max=1000"` // 视频标题
	Info  string `form:"info" json:"info" binding:"max=30000"`                 // 视频简介
	Url   string `form:"url" json:"url"`                                       // oss视频源url
	Cover string `form:"cover" json:"cover"`                                   // oss视频封面url
}

// CreateVideos
func (c *CreateVideosInfo) CreateVideos(userId string) (*models.Video, error) {
	video := &models.Video{
		Title:   c.Title,
		Info:    c.Info,
		Url:     c.Url,
		Cover:   c.Cover,
		OwnerId: userId,
	}

	// 保存数据库
	err, newVideo := models.AddVideo(video)
	return newVideo, err
}
