package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAccountController interface {
	// アカウントを登録
	Signup(ctx *gin.Context)
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
func (c *AccountController) Signup(ctx *gin.Context) {
	// リクエストボディをパース
	var json struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// アカウント登録
	err := c.service.Signup(json.Email, json.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created"})
}
