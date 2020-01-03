package services

import "github.com/zyjblockchain/gem/models"

type DeleteVideoInfo struct {
}

func (del *DeleteVideoInfo) DelVideo(videoId interface{}) error {
	return models.DeleteVideo(videoId)
}
