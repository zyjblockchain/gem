package models

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

const RankKey = "RankView"

func pageViewKey(videoId uint) string {
	return fmt.Sprintf("video:%d pageView", videoId)
}

// AddPageView通过redis记录此视频点击量
func (v *Video) AddPageView(increment int64) {
	// 增加此视频点击量
	RedisCache.IncrBy(pageViewKey(v.ID), increment)
	// 根据点击量自动排序,使用有序集合
	RedisCache.ZIncrBy(RankKey, float64(increment), strconv.Itoa(int(v.ID)))
}

// GetPageView 获取此视频的点击量
func (v *Video) GetPageView() int64 {
	num, err := RedisCache.Get(pageViewKey(v.ID)).Int64()
	if err == redis.Nil || err == nil {
		return num
	}
	panic(err)
}

// GetTopPageView 获取视频点击排名
func GetTopPageView(top int64) ([]*Video, error) {
	if top == 0 {
		return nil, errors.New("top不能为0")
	}
	videoIds, err := RedisCache.ZRevRange(RankKey, 0, top-1).Result()
	if err != nil {
		return nil, errors.New("从redis中的Rank有序集合中获取数据失败" + err.Error())
	}
	var videos = make([]*Video, 0, len(videoIds))
	for _, strId := range videoIds {
		id, _ := strconv.Atoi(strId)
		// 通过id在数据库中查询video
		err, video := GetVideo(id)
		if err != nil {
			return nil, errors.New("通过id在数据库中查询video失败：" + err.Error())
		}
		videos = append(videos, video)
	}
	return videos, nil
}
