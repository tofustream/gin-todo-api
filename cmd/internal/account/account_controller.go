package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAccountController interface {
	// アカウントを登録
	Signup(ctx *gin.Context)

	// Emailを更新
	UpdateAccountEmail(ctx *gin.Context)

	// Passwordを更新
	UpdateAccountPassword(ctx *gin.Context)

	// アカウントを削除
	DeleteAccount(ctx *gin.Context)
}

type AccountController struct {
	service IAccountApplicationService
}

func NewAccountController(service IAccountApplicationService) IAccountController {
	return &AccountController{
		service: service,
	}
}

func extractAccountID(ctx *gin.Context) (string, bool) {
	maybeAccountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return "", false
	}
	return maybeAccountID.(string), true
}

func parseJSON(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

// アカウントを登録
func (c AccountController) Signup(ctx *gin.Context) {
	var json struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if !parseJSON(ctx, &json) {
		return
	}

	err := c.service.Signup(json.Email, json.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created"})
}

// Emailを更新
func (c AccountController) UpdateAccountEmail(ctx *gin.Context) {
	accountIDStr, ok := extractAccountID(ctx)
	if !ok {
		return
	}

	var json struct {
		Email string `json:"email"`
	}
	if !parseJSON(ctx, &json) {
		return
	}

	command, err := NewUpdateAccountEmailCommand(accountIDStr, json.Email)
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

func (c AccountController) UpdateAccountPassword(ctx *gin.Context) {
	accountIDStr, ok := extractAccountID(ctx)
	if !ok {
		return
	}

	var json struct {
		Password string `json:"password"`
	}
	if !parseJSON(ctx, &json) {
		return
	}

	command, err := NewUpdateAccountPasswordCommand(accountIDStr, json.Password)
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

func (c AccountController) DeleteAccount(ctx *gin.Context) {
	accountIDStr, ok := extractAccountID(ctx)
	if !ok {
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
