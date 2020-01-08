package models

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

func InitDB(dsn string) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 设置数据库日志级别
	if gin.Mode() == gin.ReleaseMode {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	DB = db
	autoCreateTable()
}

// 自动建表
func autoCreateTable() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
}

var RedisCache *redis.Client

func InitRedis() {
	dbIndex, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       dbIndex, // 数据库index

	})

	// 检查连接情况
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	RedisCache = client
}
