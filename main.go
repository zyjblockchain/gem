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

// type User struct {
// 	Uid int32 `gorm:"primary_key"`
// 	NickName string
// 	Age int
// 	CreateAt time.Duration
// }
//
// type Order struct {
// 	Oid uint32 `gorm:"primary_key"`
// 	OrderName string
// 	CreateAt time.Duration
// 	UpdateAt time.Duration
// }

// func main() {
// 	// 连接数据库
// 	db , err := gorm.Open("mysql", "root:123456@/dbtest?charset=utf8&parseTime=True&loc=Local")
// 	defer db.Close()
// 	if err != nil {
// 		fmt.Println("err:", err)
// 	}
// 	// // 判断表是否存在
// 	// exist := db.HasTable(&User{})
// 	// exist = db.HasTable("users")
// 	// fmt.Println(exist)
// 	// // 自动建表
// 	// db.AutoMigrate(&User{},&Order{})
//
// 	// // 删除表
// 	// db.DropTableIfExists(&User{})
// 	// db.DropTableIfExists("users")
//
// 	// 删除表中的列
// 	// db.Table("orders").DropColumn("order_name")
// 	// db.Model(&Order{}).DropColumn("create_at")
//
// 	// 修改表的指定列的数据类型
// 	// db.Model(&Order{}).ModifyColumn("update_at","text")
//
// 	// CURD
// 	user := User{
// 		Uid:      3,
// 		NickName: "sandy",
// 		Age:      20,
// 	}
// 	db.Create(&user)
// 	db.Find(&User{})
//
//
//
// }
