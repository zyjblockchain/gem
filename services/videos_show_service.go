package services

import (
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
)

type ShowVideoInfo struct {
}

func (s *ShowVideoInfo) ShowVideo(videoId, userId uint) (*serializer.VideoGet, error) {
	err, video := models.GetVideo(videoId)
	if err != nil {
		return nil, err
	}
	// 判断该视频是否已经被u该user点击过了
	if !models.HasView(userId, videoId) {
		// 记录该视频被观看了一次
		models.AddPageView(video.ID, 1) // 每show一次点击量增加1
		// add该user观看记录
		models.AddRecord(userId, videoId)
	}

	videoGets, err := serializer.AssembleVideoGet([]*models.Video{video})
	if err != nil || len(videoGets) == 0 {
		return nil, err
	}
	return videoGets[0], err
}
