package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	FindAll(ctx *gin.Context)
}

type TaskController struct {
	service ITaskApplicationService
}

func NewTaskController(service ITaskApplicationService) ITaskController {
	return &TaskController{service: service}
}

func (c *TaskController) FindAll(ctx *gin.Context) {
	tasks, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
