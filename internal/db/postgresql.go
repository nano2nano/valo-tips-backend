package db

import (
	"os"

	"valo-tips/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db := getSession()
	db.AutoMigrate(&model.Tip{}, &model.Side{})
	return db
}

func getSession() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		return db
	}
}
