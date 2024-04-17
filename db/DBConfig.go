package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		db = InitDB()
		return db
	}
	return db
}

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//show sql query console
		Logger: logger.Default.LogMode(logger.Info),
	
	})
	if err != nil {
		panic(err)
	}
	return db
}
