package router

import (
	"API_BASE/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(uc *controller.UserController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		user1 := v1.Group("/users")
		{
			user1.GET("/", uc.GetAllUsers)
			user1.POST("/", uc.CreateUser)
			user1.PUT("/:id", uc.UpdateUser)
			user1.GET("/:id", uc.GetUserByID)
			user1.DELETE("/:id", uc.DeleteUserByID)

		}
	}
	return r
}
