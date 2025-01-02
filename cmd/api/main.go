package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQLドライバのインポート
	"github.com/tofustream/gin-todo-api/cmd/internal/account"
	"github.com/tofustream/gin-todo-api/cmd/internal/auth"
	"github.com/tofustream/gin-todo-api/cmd/internal/config"
	"github.com/tofustream/gin-todo-api/cmd/internal/db"
	"github.com/tofustream/gin-todo-api/cmd/internal/task"
	"github.com/tofustream/gin-todo-api/cmd/internal/user"
)

func main() {
	config.Initialize()
	database := db.SetupDB()
	defer database.Close()

	taskRepository := task.NewPostgresTaskRepository(database)
	taskService := task.NewTaskApplicationService(taskRepository)
	taskController := task.NewTaskController(taskService)

	userRepository := user.NewPostgresGeneralUserRepository(database)
	userService := user.NewGeneralUserApplicationService(userRepository)
	userController := user.NewGeneralUserController(userService)

	accountRepository := account.NewPostgresAccountRepository(database)
	accountApplicationService := account.NewAccountApplicationService(accountRepository)
	accountController := account.NewAccountController(accountApplicationService)

	authApplicationService := auth.NewAuthApplicationService(accountRepository)
	authController := auth.NewAuthController(authApplicationService)

	// Ginルーターの初期化
	r := gin.Default()
	taskRouter := r.Group(("/tasks"))
	taskRouterWithAuth := taskRouter.Group(("/"), user.AuthMiddleware(userService))
	userRouter := r.Group(("/users"))

	// タスク関連のルートを設定
	taskRouter.GET("", taskController.FindAll)
	taskRouter.GET("/:id", taskController.FindById)
	taskRouterWithAuth.POST("", taskController.Register)
	taskRouter.PUT("/:id", taskController.UpdateTaskDescription)
	taskRouter.PUT("/:id/complete", taskController.MarkTaskAsCompleted)
	taskRouter.PUT("/:id/incomplete", taskController.MarkTaskAsIncompleted)
	taskRouter.DELETE("/:id", taskController.DeleteTask)

	userRouter.GET("", userController.FindAll)
	userRouter.POST("/signup", userController.Signup)
	userRouter.POST("/login", userController.Login)

	r.POST("/signup", accountController.Signup)
	r.POST("/login", authController.Login)

	// サーバーをポート8080で起動
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
