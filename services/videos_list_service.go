package services

import "github.com/zyjblockchain/gem/models"

type BatchVideosInfo struct {
	Start int `form:"start" json:"start"`
	Limit int `form:"limit" json:"limit"`
}

func (l *BatchVideosInfo) BatchVideos() ([]*models.Video, error) {
	return models.GetBatchVideos(l.Start, l.Limit)
}
