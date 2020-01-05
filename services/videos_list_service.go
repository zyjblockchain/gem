package services

import (
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
)

type BatchVideosInfo struct {
	Start int `form:"start" json:"start"`
	Limit int `form:"limit" json:"limit"`
}

func (l *BatchVideosInfo) BatchVideos() ([]*serializer.VideoGet, error) {
	videos, err := models.GetBatchVideos(l.Start, l.Limit)
	if err != nil {
		return nil, err
	}
	return serializer.AssembleVideoGet(videos)
}
