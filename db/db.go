package db

import (
	"go-gin-crud-sample/models"
	"go-gin-crud-sample/pkg/setting"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dsnStrings := []string{
		setting.DatabaseSetting.User,
		":",
		setting.DatabaseSetting.Password,
		"@tcp(",
		setting.DatabaseSetting.Host,
		")/",
		setting.DatabaseSetting.DatabaseName,
		"?charset=utf8mb4&parseTime=True&loc=Local",
	}
	dsn := strings.Join(dsnStrings, "")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	sqldb, err := db.DB()
	if err != nil {
		return
	}
	if err := sqldb.Close(); err != nil {
		panic(err)
	}
}
