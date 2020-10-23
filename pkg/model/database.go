package model

import (
	"github.com/general252/gout/ulog"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

// InitDataBase 初始化数据库
func InitDataBase(dsn string) error {
	//dbName := fmt.Sprintf("%v.db", uapp.GetExeName())
	// sqlite.Open(dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		ulog.Error("gorm open fail: %v", err)
		return err
	}

	db.Logger.LogMode(logger.Info)

	err = db.AutoMigrate(
		new(MUser),
	)

	if err != nil {
		ulog.Error("gorm AutoMigrate: %v", err)
		return err
	}

	return nil
}

// GetDataBase 数据库
func GetDataBase() *gorm.DB {
	return db
}
