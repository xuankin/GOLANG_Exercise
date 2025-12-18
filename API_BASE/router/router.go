package router

import (
	"API_BASE/config"
	"API_BASE/controller"
	"API_BASE/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(uc *controller.UserController, conf *config.Config) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", uc.CreateUser)
			auth.POST("/login", uc.Login)
			auth.POST("/refresh", uc.RefreshToken)
		}
		userRoutes := v1.Group("/users")
		userRoutes.Use(middleware.AuthMiddleWare(conf.JWTSecret))
		{
			userRoutes.GET("/", uc.GetAllUsers)
			userRoutes.POST("/", uc.CreateUser)
			userRoutes.GET("/:id", uc.GetUserByID)
			userRoutes.PUT("/:id", uc.UpdateUser)
			userRoutes.DELETE("/:id", uc.DeleteUserByID)
			userRoutes.POST("/logout", uc.Logout)
		}
	}
	return r
}
