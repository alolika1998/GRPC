package main

import (
	"log"

	"github.com/go/grpc/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/add/:a/:b", handlers.AdditionHandler)
	g.GET("/sub/:a/:b", handlers.SubtractionHandler)
	g.GET("/mult/:a/:b", handlers.MultiplicationHandler)
	g.GET("/div/:a/:b", handlers.DivisionHandler)
	g.GET("/rem/:a/:b", handlers.RemainderHandler)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
