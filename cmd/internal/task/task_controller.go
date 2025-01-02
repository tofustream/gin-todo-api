package task

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITaskController interface {
	FindAllByAccountID(ctx *gin.Context)
	FindById(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTaskDescription(ctx *gin.Context)
	MarkTaskAsCompleted(ctx *gin.Context)
	MarkTaskAsIncompleted(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

type TaskController struct {
	service ITaskApplicationService
}

func NewTaskController(service ITaskApplicationService) ITaskController {
	return &TaskController{service: service}
}

func (c *TaskController) FindAllByAccountID(ctx *gin.Context) {
	maybeAccountID := ctx.MustGet("accountID")
	accountIDStr, ok := maybeAccountID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}
	accountID, err := uuid.Parse(accountIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}
	dtos, err := c.service.FindAllByAccountID(accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": dtos})
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

func (c *TaskController) CreateTask(ctx *gin.Context) {
	maybeAccountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}
	accountIDStr := maybeAccountID.(string)
	log.Printf("accountID: %s", accountIDStr)

	var json struct {
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.CreateTask(json.Description, accountIDStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

func (c *TaskController) UpdateTaskDescription(ctx *gin.Context) {
	paramID := ctx.Param("id")
	var json struct {
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command, err := NewUpdateTaskDescriptionCommand(paramID, json.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dto, err := c.service.Update(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": dto})
}

func (c *TaskController) MarkTaskAsCompleted(ctx *gin.Context) {
	paramID := ctx.Param("id")

	command, err := NewMarkTaskAsCompleteCommand(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dto, err := c.service.Update(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": dto})
}

func (c *TaskController) MarkTaskAsIncompleted(ctx *gin.Context) {
	paramID := ctx.Param("id")

	command, err := NewMarkTaskAsIncompleteCommand(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dto, err := c.service.Update(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": dto})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	paramID := ctx.Param("id")

	command, err := NewMarkAsDeletedCommand(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dto, err := c.service.Update(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": dto})
}
