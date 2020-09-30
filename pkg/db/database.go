package db


import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func InitDataBase() error {
	var err error
	db, err = gorm.Open("sqlite3", "plugin.db")
	if err != nil {
		return err
	}

	db.LogMode(true)

	db.AutoMigrate(new(MPlugin))
	db.AutoMigrate(new(MPluginVersion))

	return nil
}

func GetDataBase() *gorm.DB {
	return db
}
