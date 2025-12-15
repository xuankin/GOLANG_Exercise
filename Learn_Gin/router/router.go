package router

import (
	"Learn_Gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userController *controller.UserController) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUserById)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
