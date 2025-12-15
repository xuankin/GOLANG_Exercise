package main

import (
	"Learn_Gin/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
