package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tofustream/gin-todo-api/cmd/internal/config"
	"github.com/tofustream/gin-todo-api/cmd/internal/db"
)

// マイグレーションファイルのパス
var migrationFiles = []string{
	"./cmd/internal/migrations/20250101_create_general_users_table.sql",
	"./cmd/internal/migrations/20250101_create_tasks_table.sql",
}

func main() {
	// 環境変数の初期化
	config.Initialize()

	// データベース接続のセットアップ
	database := db.SetupDB()
	defer database.Close()

	for _, migrationFile := range migrationFiles {
		// マイグレーションファイルの内容を読み込む
		migrationSQL, err := os.ReadFile(migrationFile)
		if err != nil {
			log.Panicf("Error reading migration file %s: %v", migrationFile, err)
		}

		// マイグレーションSQLを実行
		_, err = database.Exec(string(migrationSQL))
		if err != nil {
			log.Panicf("Error executing migration from file %s: %v", migrationFile, err)
		}

		// 成功メッセージ
		fmt.Printf("Migration from file %s executed successfully\n", migrationFile)
	}
}
