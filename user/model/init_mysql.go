package model

import (
	"fmt"
	conf "user/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// dsn := "root:123456@tcp(192.168.0.6:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.MysqlUsername, conf.MysqlPassword, conf.MysqlIp, conf.MysqlPort, conf.MysqlDatabase)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		//SkipDefaultTransaction: true, //禁用事务
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}

	// 迁移数据库
	migrate()
}
