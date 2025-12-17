package main

import (
	"API_BASE/config"
	"API_BASE/controller"
	"API_BASE/entity"
	"API_BASE/repository"
	"API_BASE/router"
	"API_BASE/service"
	"log"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Khong the load config file.")
	}
	db := config.ConnectDB(&conf)
	db.AutoMigrate(&entity.User{})
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	r := router.SetupRouter(userController)
	log.Println("Server start...", conf.ServerAddr)
	if err := r.Run(conf.ServerAddr); err != nil {
		log.Fatal("Server start error.")
	}
}
