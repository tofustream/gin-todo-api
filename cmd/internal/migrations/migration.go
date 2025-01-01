package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tofustream/gin-todo-api/cmd/internal/config"
	"github.com/tofustream/gin-todo-api/cmd/internal/db"
)

// マイグレーションファイルのパス
const migrationFile string = "./cmd/internal/migrations/20250101_create_tasks_table.sql"

func main() {
	// 環境変数の初期化
	config.Initialize()

	// データベース接続のセットアップ
	database := db.SetUpDB()
	defer database.Close()

	// マイグレーションファイルの内容を読み込む
	migrationSQL, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Panicf("Error reading migration file: %v", err)
	}

	// マイグレーションSQLを実行
	_, err = database.Exec(string(migrationSQL))
	if err != nil {
		log.Panicf("Error executing migration: %v", err)
	}

	// 成功メッセージ
	fmt.Println("Migration executed successfully")
}
