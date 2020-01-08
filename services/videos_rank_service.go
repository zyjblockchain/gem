package services

import (
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
)

type RankInfo struct {
}

func (r *RankInfo) GetRankVideos(top int64) ([]*serializer.VideoGet, error) {
	videos, err := models.GetTopPageView(top)
	if err != nil {
		return nil, err
	}
	return serializer.AssembleVideoGet(videos)
}
