package handles

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/serializer"
	"github.com/zyjblockchain/gem/services"
)

// Register 用户注册接口
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service services.UserRegisterInfo
		// 解析传入的参数
		if err := c.ShouldBind(&service); err == nil {
			// 解析成功则进行注册
			if user, err := service.RegisterUser(); err != nil {
				c.JSON(200, &serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "注册失败",
					Error:  err.Error(),
				})
			} else {
				// 返回注册的user存储信息
				c.JSON(200, &serializer.Response{
					Status: 20000,
					Data:   user,
					Msg:    "注册成功",
					Error:  "",
				})
			}
		} else {
			c.JSON(400, &serializer.Response{
				Status: 400001,
				Data:   nil,
				Msg:    "注册参数解析失败",
				Error:  err.Error(),
			})
		}
	}
}

// Login 用户登录接口
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginInfo services.UserLoginInfo
		if err := c.ShouldBind(&loginInfo); err == nil {
			user, err := loginInfo.UserLogin()
			if err != nil {
				// 用户登录失败
				c.JSON(200, &serializer.Response{
					Status: 40001,
					Data:   nil,
					Msg:    "登录不成功",
					Error:  err.Error(),
				})
			} else {
				// 登录校验成功
				// 设置session
				s := sessions.Default(c)
				// 删除之前设置的kv
				s.Clear()
				// 重新设置kv
				s.Set("user_id", user.ID)
				// 保存当前请求的session
				_ = s.Save()

				// 返回登录结果
				c.JSON(200, serializer.Response{
					Status: 200,
					Data:   *user,
					Msg:    "登录成功",
					Error:  "",
				})
			}
		} else {
			// 解析请求参数出错
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "解析请求参数出错",
				Error:  err.Error(),
			})
		}
	}
}

// 拉取自己的用户信息
func GetMine() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, exist := c.Get("user"); exist {
			if u, ok := user.(*models.User); ok {
				c.JSON(200, serializer.Response{
					Status: 200,
					Data:   *u,
					Msg:    "得到用户信息",
					Error:  "",
				})
			}
		} else {
			// 等不到
			c.JSON(200, serializer.Response{
				Status: 40001,
				Data:   nil,
				Msg:    "获取当前用户信息失败，可能没有登录",
				Error:  "not found",
			})
		}
	}
}

// 登出
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 删除session记录
		s := sessions.Default(c)
		s.Clear()
		_ = s.Save()

		c.JSON(200, serializer.Response{
			Status: 200,
			Data:   nil,
			Msg:    "登出成功",
			Error:  "",
		})
	}
}
