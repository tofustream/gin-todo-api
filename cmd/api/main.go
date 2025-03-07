package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQLドライバのインポート
	"github.com/tofustream/gin-todo-api/cmd/internal/account"
	"github.com/tofustream/gin-todo-api/cmd/internal/auth"
	"github.com/tofustream/gin-todo-api/cmd/internal/config"
	"github.com/tofustream/gin-todo-api/cmd/internal/db"
	"github.com/tofustream/gin-todo-api/cmd/internal/task"
)

func main() {
	config.Initialize()
	database := db.SetupDB()
	defer database.Close()

	taskRepository := task.NewPostgresTaskRepository(database)
	taskService := task.NewTaskApplicationService(taskRepository)
	taskController := task.NewTaskController(taskService)

	accountRepository := account.NewPostgresAccountRepository(database)
	accountApplicationService := account.NewAccountApplicationService(accountRepository)
	accountController := account.NewAccountController(accountApplicationService)

	authApplicationService := auth.NewAuthApplicationService(accountRepository)
	authController := auth.NewAuthController(authApplicationService)

	// Ginルーターの初期化
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/signup", accountController.Signup)
	r.POST("/login", authController.Login)

	// ミドルウェアで使用するシークレットキーを環境変数から取得
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("Environment variable SECRET_KEY is not set")
	}

	accountRouter := r.Group(("/accounts"))
	accountRouterWithAuth := accountRouter.Group(("/"), auth.AuthMiddleware(secretKey))
	accountRouterWithAuth.GET("", accountController.FindAccount)
	accountRouterWithAuth.PATCH("", accountController.UpdateAccount)
	accountRouterWithAuth.DELETE("", accountController.DeleteAccount)

	taskRouter := r.Group(("/tasks"))
	taskRouterWithAuth := taskRouter.Group(("/"), auth.AuthMiddleware(secretKey))
	taskRouterWithAuth.GET("/:id", taskController.FindTask)
	taskRouterWithAuth.GET("", taskController.FindAllTasksByAccountID)
	taskRouterWithAuth.POST("", taskController.CreateTask)
	taskRouterWithAuth.PATCH("/:id", taskController.UpdateTask)
	taskRouterWithAuth.DELETE("/:id", taskController.DeleteTask)

	// サーバーをポート8080で起動
	err := r.Run(":8080")
	if err != nil {
		log.Panicf("Failed to start server: %v", err)
	}
}
