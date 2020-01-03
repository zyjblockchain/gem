package services

import "github.com/zyjblockchain/gem/models"

type OwnerVideos struct {
}

func (v *OwnerVideos) GetVideos(userId interface{}) ([]*models.Video, error) {
	return models.GetVideosByUserId(userId)
}
