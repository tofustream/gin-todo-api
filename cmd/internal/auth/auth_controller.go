package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(ctx *gin.Context)
}

type AuthController struct {
	service IAuthApplicationService
}

func NewAuthController(service IAuthApplicationService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var json struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(json.Email, json.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": *token})
}
