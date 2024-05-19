package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadConfig() string {
	// .envファイルのパスを動的に切り替える
	envPath := ".env"
	if os.Getenv("RUN_ENV") == "DOCKER" {
		envPath = "/app/.env"
	}

	// .envファイルを読み込む
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Error loading .env file from path: %s", envPath)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT_NUMBER"),
		os.Getenv("USER_NAME"),
		os.Getenv("DBNAME"),
		os.Getenv("PASSWORD"),
	)
	return dsn
}
