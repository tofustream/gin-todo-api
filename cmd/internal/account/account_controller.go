package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tofustream/gin-todo-api/cmd/internal/auth/authhelper"
)

type IAccountController interface {
	// アカウントを登録
	Signup(ctx *gin.Context)

	// アカウントを更新
	UpdateAccount(ctx *gin.Context)

	// アカウントを削除
	DeleteAccount(ctx *gin.Context)

	// アカウントを取得
	FindAccount(ctx *gin.Context)
}

type AccountController struct {
	service IAccountApplicationService
}

func NewAccountController(service IAccountApplicationService) IAccountController {
	return &AccountController{
		service: service,
	}
}

// アカウントを登録
func (c AccountController) Signup(ctx *gin.Context) {
	var json struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.Signup(json.Email, json.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created"})
}

func (c AccountController) UpdateAccount(ctx *gin.Context) {
	accountIDStr, exists := authhelper.GetAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}

	var json struct {
		Email    *string `json:"email,omitempty"`
		Password *string `json:"password,omitempty"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isUpdated := false
	if json.Email != nil {
		command, err := NewUpdateAccountEmailCommand(accountIDStr, *json.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err = c.service.UpdateAccount(command)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		isUpdated = true
	}
	if json.Password != nil {
		command, err := NewUpdateAccountPasswordCommand(accountIDStr, *json.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err = c.service.UpdateAccount(command)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		isUpdated = true
	}

	// 更新されなかった場合
	if !isUpdated {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
		return
	}

	accountDTO, err := c.service.FindAccount(accountIDStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": accountDTO})
}

func (c AccountController) DeleteAccount(ctx *gin.Context) {
	accountIDStr, exists := authhelper.GetAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}

	command, err := NewMarkAsDeletedCommand(accountIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := c.service.UpdateAccount(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": dto})
}

func (c AccountController) FindAccount(ctx *gin.Context) {
	accountIDStr, exists := authhelper.GetAccountIDFromContext(ctx)
	if !exists {
		ctx.AbortWithStatus((http.StatusUnauthorized))
		return
	}

	dto, err := c.service.FindAccount(accountIDStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": dto})
}
