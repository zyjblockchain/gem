package services

import (
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
)

type ShowVideoInfo struct {
}

func (s *ShowVideoInfo) ShowVideo(videoId interface{}) (*serializer.VideoGet, error) {
	err, video := models.GetVideo(videoId)
	if err != nil {
		return nil, err
	}

	// 记录该视频被观看了一次
	video.AddPageView(1) // 每show一次点击量增加1

	videoGets, err := serializer.AssembleVideoGet([]*models.Video{video})
	if err != nil || len(videoGets) == 0 {
		return nil, err
	}
	return videoGets[0], err
}
