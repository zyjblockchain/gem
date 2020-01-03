package services

import "github.com/zyjblockchain/gem/models"

type ShowVideoInfo struct {
}

func (s *ShowVideoInfo) ShowVideo(videoId interface{}) (*models.Video, error) {
	err, video := models.GetVideo(videoId)
	return video, err
}
