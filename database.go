package main

import (
	// "github.com/eternaleight/go-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// ユーザーモデルを自動マイグレーション
	// db.AutoMigrate(&models.User{}, &models.Post{}, &models.Profile{}, &models.Product{}, &models.Purchase{}, &models.Image{}, &models.Favorite{}, &models.RequestCard{})
	return db
}
