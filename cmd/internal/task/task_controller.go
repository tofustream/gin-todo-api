package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Add(ctx *gin.Context)
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

func (c *TaskController) FindById(ctx *gin.Context) {
	paramID := ctx.Param("id")

	task, err := c.service.FindById(paramID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": task})
}

func (c *TaskController) Add(ctx *gin.Context) {
	var json struct {
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := c.service.Add(json.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"task": task})
}
