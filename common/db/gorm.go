package db

import (
	"fmt"
	"log"
	"newe-serve/common/nelog"
	"newe-serve/common/setting"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	Db *gorm.DB
)

type logprint struct {
}

func (a logprint) Printf(s string, v ...interface{}) {
	nelog.TRACE.Printf(s, v...)
}

func Setup() {

	newLogger := logger.New(
		logprint{}, // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", setting.SqlDb.User, setting.SqlDb.Password, setting.SqlDb.Host, setting.SqlDb.Name)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			//设置全局禁用表自动复数
			SingularTable: true,
		},
		SkipDefaultTransaction: true, //禁用默认事务
	})
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	Db = db.Debug()
}

// func CloseDB() {
// 	defer Db.Close()
// }

// func Verifyform(tablename string, id int64, where map[string]interface{}) bool {
// 	var count int64

// 	if id > 0 {
// 		Db.Table(tablename).Where(where).Where("id <> ?", id).Count(&count)

// 	} else {
// 		Db.Table(tablename).Where(where).Count(&count)
// 	}

// 	fmt.Println("条数：i%", count)
// 	if count > 0 {
// 		return true
// 	}
// 	return false
// }
