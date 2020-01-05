package services

import (
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
)

type OwnerVideos struct {
}

func (v *OwnerVideos) GetVideos(userId interface{}) ([]*serializer.VideoGet, error) {
	videos, err := models.GetVideosByUserId(userId)
	if err != nil {
		return nil, err
	}
	return serializer.AssembleVideoGet(videos)
}
