package conf

import (
	"github.com/joho/godotenv"
	"github.com/zyjblockchain/gem/models"
	"github.com/zyjblockchain/gem/services/oss"
	"os"
)

// Init 初始化数据库
func Init() {
	// 从本地的.env文件中读取配置文件到环境变量中
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// 初始化bucket pool
	oss.Pool = oss.NewBucketPool()
	// 链接数据库
	models.InitDB(os.Getenv("MYSQL_DSN"))

	// 连接redis
	models.InitRedis()
}
