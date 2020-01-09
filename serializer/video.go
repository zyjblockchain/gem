package serializer

import (
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/services/oss"
)

type VideoPut struct {
	models.Video
	VideoPutUrl  string `json:"videoPutUrl"`
	AvatarPutUrl string `json:"avatarPutUrl"`
}

func AssembleVideoPut(videos []*models.Video) ([]*VideoPut, error) {
	var result = make([]*VideoPut, 0, len(videos))
	for _, video := range videos {
		// 生成put url
		videoUrl, err := oss.PutSignedUrl(oss.Pool.GetBucket(), video.VideoPath)
		if err != nil {
			return nil, err
		}
		avatarUrl, err := oss.PutSignedUrl(oss.Pool.GetBucket(), video.AvatarPath)
		if err != nil {
			return nil, err
		}
		vp := &VideoPut{
			Video:        *video,
			VideoPutUrl:  videoUrl,
			AvatarPutUrl: avatarUrl,
		}
		result = append(result, vp)
	}
	return result, nil
}

type VideoGet struct {
	models.Video
	VideoHits    int64  `json:"videoHits"` // 视频点击量
	VideoGetUrl  string `json:"videoGetUrl"`
	AvatarGetUrl string `json:"avatarGetUrl"`
}

func AssembleVideoGet(videos []*models.Video) ([]*VideoGet, error) {
	var result = make([]*VideoGet, 0, len(videos))
	for _, video := range videos {
		// 获取get url
		videoUrl, err := oss.GetSignedUrl(oss.Pool.GetBucket(), video.VideoPath)
		if err != nil {
			return nil, err
		}
		avatarUrl, err := oss.GetSignedUrl(oss.Pool.GetBucket(), video.AvatarPath)
		if err != nil {
			return nil, err
		}

		vg := &VideoGet{
			Video:        *video,
			VideoHits:    models.GetPageView(video.ID),
			VideoGetUrl:  videoUrl,
			AvatarGetUrl: avatarUrl,
		}
		result = append(result, vg)
	}
	return result, nil
}
