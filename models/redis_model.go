package models

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

const RankKey = "RankView"

func pageViewKey(videoId uint) string {
	return fmt.Sprintf("videoId:%d", videoId)
}

func UserKey(userId uint) string {
	return fmt.Sprintf("userId: %d", userId)
}

// AddPageView通过redis记录此视频点击量
func AddPageView(videoId uint, increment int64) {
	// 增加此视频点击量
	RedisCache.IncrBy(pageViewKey(videoId), increment)
	// 根据点击量自动排序,使用有序集合
	RedisCache.ZIncrBy(RankKey, float64(increment), strconv.Itoa(int(videoId)))
}

// GetPageView 获取此视频的点击量
func GetPageView(videoId uint) int64 {
	num, err := RedisCache.Get(pageViewKey(videoId)).Int64()
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

// AddRecord 记录每个用户的观看记录
func AddRecord(userId, videoId uint) {
	err := RedisCache.SAdd(UserKey(userId), pageViewKey(videoId)).Err()
	if err != nil {
		panic(err)
	}
}

// HasView查看此用户是否已经看过此视频
func HasView(userId, videoId uint) bool {
	ok, err := RedisCache.SIsMember(UserKey(userId), pageViewKey(videoId)).Result()
	if err != nil {
		panic(err)
	}
	return ok
}
