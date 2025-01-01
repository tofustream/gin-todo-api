package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQLドライバのインポート
	"github.com/tofustream/gin-todo-api/cmd/internal/config"
	"github.com/tofustream/gin-todo-api/cmd/internal/db"
	"github.com/tofustream/gin-todo-api/cmd/internal/task"
)

func main() {
	config.Initialize()
	database := db.SetUpDB()
	defer database.Close()

	taskRepository := task.NewPostgresTaskRepository(database)
	taskService := task.NewTaskApplicationService(taskRepository)
	taskController := task.NewTaskController(taskService)

	// Ginルーターの初期化
	r := gin.Default()

	// タスク関連のルートを設定
	r.GET("/tasks", taskController.FindAll)
	r.GET("/tasks/:id", taskController.FindById)
	r.POST("/tasks", taskController.Register)
	r.PUT("/tasks/:id", taskController.UpdateTaskDescription)
	r.PUT("/tasks/:id/complete", taskController.MarkTaskAsCompleted)
	r.PUT("/tasks/:id/incomplete", taskController.MarkTaskAsIncompleted)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	// サーバーをポート8080で起動
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
