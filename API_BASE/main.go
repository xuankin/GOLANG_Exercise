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
		log.Fatal("Cannot load config file")
	}
	db := config.ConnectDB(&conf)
	rdb := config.ConnectRedis(&conf)
	db.AutoMigrate(&entity.User{})
	userRepo := repository.NewUserRepository(db, rdb)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService, &conf)
	r := router.SetupRouter(userController, &conf)
	log.Println("Server start...", conf.ServerAddr)
	if err := r.Run(conf.ServerAddr); err != nil {
		log.Fatal("Server start error.")
	}
}
