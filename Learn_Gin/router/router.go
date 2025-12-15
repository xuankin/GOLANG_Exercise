package router

import (
	"Learn_Gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	usersController := controller.NewUserController()
	r.GET("/users", usersController.GetUsers)
}
