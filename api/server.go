package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func startApiServer() {

	if ctx.Conf().API.Disabled {
		return
	}

	engine := gin.New()
	engine.GET("/status", status)

	api := engine.Group("/api")
	api.GET("/", status)

	terminals := api.Group("/terminal")
	terminals.GET("/ws", terminalWebSocket)

	log.Printf("Starting API server on %s", ":8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Error while starting API server: %s", err.Error())
	}
}
