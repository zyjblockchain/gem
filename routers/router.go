package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zyjblockchain/gem/handles"
	"github.com/zyjblockchain/gem/middleware"
	"os"
)

// 全局路由设置
func NewRouter(addr string) {
	r := gin.Default()

	// 执行中间件
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.SetLoginUser())

	// 注册路由
	v1 := r.Group("/api/v1")
	{
		// 心跳检测接口
		v1.POST("ping", handles.Ping)

		// 1. 用户注册接口
		v1.POST("user/register", handles.Register())
		// 2. 用户登录接口
		v1.POST("user/login", handles.Login())

		v1.GET("video/:id", handles.ShowVideo()) // 点播

		// 3. 需要登录保护
		authed := v1.Group("/")
		// 需要登录授权才能访问的接口
		authed.Use(middleware.AuthRequired())
		{
			// 拉取自己的用户信息
			authed.GET("user/me", handles.GetMine())
			// 登出
			authed.DELETE("user/logout", handles.Logout())
		}

		// 4. 视频相关接口
		v1.GET("videos", handles.GetOwnerVideos())    // 获取自己上传的videos
		v1.POST("videos", handles.CreateVideo())      // 视频投K稿
		v1.POST("batchVideos", handles.ListVideo())   // 视频列表
		v1.PUT("video/:id", handles.UpdateVideo())    // 更新视频内容
		v1.DELETE("video/:id", handles.DeleteVideo()) // 删除视频

		// 5. 拉取每日视频点击量排行
		v1.GET("rank/:top", handles.Rank())
	}

	if err := r.Run(addr); err != nil {
		panic(err)
	}
}
