package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IGeneralUserController interface {
	FindAll(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type GeneralUserController struct {
	service IGeneralUserApplicationService
}

func NewGeneralUserController(service IGeneralUserApplicationService) IGeneralUserController {
	return &GeneralUserController{
		service: service,
	}
}

func (c *GeneralUserController) FindAll(ctx *gin.Context) {
	users, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *GeneralUserController) Register(ctx *gin.Context) {
	var json struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.Register(json.Email, json.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User has been registered"})
}
