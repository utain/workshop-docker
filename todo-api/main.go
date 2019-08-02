package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting...")
	serv := gin.New()
	serv.Use(gin.Recovery())
	TodoRouter(serv)
	serv.Run(":4000")
}
