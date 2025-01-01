package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	// .envファイルを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
