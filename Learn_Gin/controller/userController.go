package controller

import (
	"Learn_Gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
}
func (uc *UserController) GetUsers(c *gin.Context) {
	result := uc.userService.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
