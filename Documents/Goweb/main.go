package main

import (
	"Goweb/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	p := router.Group("/products")

	router.GET("/ping", handlers.Ping)
	p.GET("/", handlers.List)

	router.Run()
}
