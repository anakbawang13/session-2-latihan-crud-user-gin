package main

import (
	"gin/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SetupRouter(r)

	log.Println("Running server on port 8080")
	r.Run(":8080")
}
