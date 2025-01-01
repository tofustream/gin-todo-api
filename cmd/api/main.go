package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/cmd/internal/task"
)

func main() {
	uuid1, _ := uuid.Parse("eda07994-8d54-4a7c-9757-b0f8ea0ba736")
	uuid2, _ := uuid.Parse("f3b3b3b3-4b3b-4b3b-4b3b-4b3b4b3b4b3b")
	taskID1, _ := task.NewTaskID(uuid1)
	taskID2, _ := task.NewTaskID(uuid2)
	description1, _ := task.NewTaskDescription("Task 1")
	description2, _ := task.NewTaskDescription("Task 2")
	tasks := make(map[task.TaskID]task.Task)
	tasks[taskID1] = task.NewTask(taskID1, description1)
	tasks[taskID2] = task.NewTask(taskID2, description2)
	taskRepository := task.NewInMemoryTaskRepository(tasks)
	taskService := task.NewTaskApplicationService(taskRepository)
	taskController := task.NewTaskController(taskService)

	r := gin.Default()
	r.GET("/tasks", taskController.FindAll)
	r.GET("/tasks/:id", taskController.FindById)
	r.POST("/tasks", taskController.Register)
	r.PUT("/tasks/:id", taskController.UpdateTaskDescription)
	r.PUT("/tasks/:id/complete", taskController.MarkTaskAsComplete)
	r.PUT("/tasks/:id/incomplete", taskController.MarkTaskAsIncompleteCommand)
	err := r.Run()
	if err != nil {
		log.Println(err)
	}
}
