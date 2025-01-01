package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func SetUpDB() *sql.DB {
	// データベース接続情報を環境変数から取得
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// データベース接続を初期化
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf("failed to open database: %v", err)
	}

	// 実際にデータベースへ接続できるか確認
	err = db.Ping()
	if err != nil {
		log.Panicf("failed to connect to the database: %v", err)
	}

	return db
}
