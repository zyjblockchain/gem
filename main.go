package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zyjblockchain/gem/conf"
	"github.com/zyjblockchain/gem/routers"
)

func main() {
	// 初始化数据库
	conf.Init()
	routers.NewRouter(":3000")
}
