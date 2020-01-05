package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Title      string // 视频标题
	Info       string // 视频简介
	VideoPath  string // 视频源oss路径
	AvatarPath string // 视频封面oss路径
	OwnerId    string // 视频上传者user id
}

func AddVideo(v *Video) (error, *Video) {
	err := DB.Create(v).Error
	return err, v
}

func GetVideo(videoId interface{}) (error, *Video) {
	var video Video
	err := DB.First(&video, videoId).Error
	return err, &video
}

// 通过userId获取所有相关的视频列表
func GetVideosByUserId(userId interface{}) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("owner_id = ?", userId).Find(&videos).Error
	return videos, err
}

// GetBatchVideos
func GetBatchVideos(start, limit int) ([]*Video, error) {
	var videos []*Video
	var total int
	err := DB.Model(&Video{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	if start > total {
		return nil, errors.New("start 超过了总数量")
	}

	err = DB.Limit(limit).Offset(start).Find(&videos).Error
	return videos, err
}

// UpdateVideo
func UpdateVideo(videoId interface{}, title, info string) (*Video, error) {
	var video Video
	err := DB.First(&video, videoId).Error
	if err != nil {
		return nil, err
	}
	video.Title = title
	video.Info = info
	// save
	err = DB.Save(&video).Error
	return &video, err
}

// delete 软删除
func DeleteVideo(videoId interface{}) error {
	var v Video
	err := DB.First(&v, videoId).Error
	if err != nil {
		return err
	}
	return DB.Delete(&v).Error
}
