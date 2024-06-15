package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitMysql() *gorm.DB {
	//if global.Config.Mysql.Host == "" {
	//	log.Println("mysql config is empty")
	//	return nil
	//}
	//dsn := global.Config.Mysql.Dsn()
	dsn := "root:root@tcp(127.0.0.1:3306)/im_server_db?charset=utf8mb4&parseTime=True&loc=Local"
	var mysqlLogger logger.Interface
	//if global.Config.System.Env == "debug" {
	//	// 开发环境显示所有的sql
	//	mysqlLogger = logger.Default.LogMode(logger.Info)
	//} else {
	//	mysqlLogger = logger.Default.LogMode(logger.Error) // 生产环境只显示错误sql
	//}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("[%s] mysql connect error", dsn))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxOpenConns(100)              // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // SetConnMaxLifetime 设置了连接可复用的最大时间。不能超过mysql的wait_timeout设置，否则会出现mysql has gone away错误
	return db
}
