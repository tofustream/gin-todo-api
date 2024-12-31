package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/cmd/internal/task"
)

func main() {
	uuid1, _ := uuid.NewRandom()
	uuid2, _ := uuid.NewRandom()
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
	r.POST("/tasks", taskController.Add)
	err := r.Run()
	if err != nil {
		log.Println(err)
	}
}
