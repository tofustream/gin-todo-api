package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const accountIDKey = "accountID"

type ITaskController interface {
	FindAllByAccountID(ctx *gin.Context)
	FindTask(ctx *gin.Context)
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

// context から account id を取得する
func getAccountIDFromContext(ctx *gin.Context) (string, bool) {
	maybeAccountID, exists := ctx.Get(accountIDKey)
	if !exists {
		return "", false
	}
	return maybeAccountID.(string), true
}

// account id に紐付くすべての task を取得する
func (c TaskController) FindAllByAccountID(ctx *gin.Context) {
	accountIDStr, exists := getAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}
	dtos, err := c.service.FindAllByAccountID(accountIDStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": dtos})
}

// task id と account id から task を取得する
func (c TaskController) FindTask(ctx *gin.Context) {
	accountIDStr, exists := getAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}

	taskIDStr := ctx.Param("id")
	task, err := c.service.FindTask(taskIDStr, accountIDStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": task})
}

// task を作成する
func (c TaskController) CreateTask(ctx *gin.Context) {
	accountIDStr, exists := getAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}

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

// task の description を更新する
func (c TaskController) UpdateTaskDescription(ctx *gin.Context) {
	accountIDStr, exists := getAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}

	taskIDStr := ctx.Param("id")
	var json struct {
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command, err := NewUpdateTaskDescriptionCommand(taskIDStr, json.Description, accountIDStr)
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

func (c TaskController) MarkTaskAsCompleted(ctx *gin.Context) {
	accountIDStr, exists := getAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}
	taskIDStr := ctx.Param("id")
	command, err := NewMarkTaskAsCompleteCommand(taskIDStr, accountIDStr)
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

func (c TaskController) MarkTaskAsIncompleted(ctx *gin.Context) {
	accountIDStr, exists := getAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}
	taskIDStr := ctx.Param("id")
	command, err := NewMarkTaskAsIncompleteCommand(taskIDStr, accountIDStr)
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

func (c TaskController) DeleteTask(ctx *gin.Context) {
	accountIDStr, exists := getAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}
	taskIDStr := ctx.Param("id")
	command, err := NewMarkAsDeletedCommand(taskIDStr, accountIDStr)
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
