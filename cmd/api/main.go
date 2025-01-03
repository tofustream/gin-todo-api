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

	taskRouter := r.Group(("/tasks"))
	taskRouterWithAuth := taskRouter.Group(("/"), auth.AuthMiddleware(os.Getenv("SECRET_KEY")))
	taskRouterWithAuth.GET("/:id", taskController.FindTask)
	taskRouterWithAuth.GET("", taskController.FindAllByAccountID)
	taskRouterWithAuth.POST("", taskController.CreateTask)
	taskRouterWithAuth.PUT("/:id", taskController.UpdateTaskDescription)
	taskRouterWithAuth.PUT("/:id/complete", taskController.MarkTaskAsCompleted)
	taskRouterWithAuth.PUT("/:id/incomplete", taskController.MarkTaskAsIncompleted)
	taskRouterWithAuth.DELETE("/:id", taskController.DeleteTask)

	// サーバーをポート8080で起動
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
