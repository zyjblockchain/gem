package services

import "github.com/zyjblockchain/gem/models"

type UpdateVideoInfo struct {
	Title string `form:"title" json:"title" binding:"min=3,max=1000"`
	Info  string `form:"info" json:"info" binding:"max=30000"`
}

func (up *UpdateVideoInfo) UpdateVideo(videoId string) (*models.Video, error) {
	return models.UpdateVideo(videoId, up.Title, up.Info)
}
