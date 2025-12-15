package main

import (
	"Learn_Gin/controller"
	"Learn_Gin/repository"
	"Learn_Gin/router"
	"Learn_Gin/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	userRepos := repository.NewUserRepository("user.json")
	userService := service.NewUserService(userRepos)
	userController := controller.NewUserController(userService)
	r := gin.Default()
	router.SetupRouter(r, userController)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
