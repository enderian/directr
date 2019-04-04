package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func startApiServer() {

	if ctx.Conf().API.Disabled {
		return
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.GET("/status", status)

	api := engine.Group("/api")
	usersGroup(api.Group("/users"))
	terminalsGroup(api.Group("/terminal"))

	log.Printf("Starting API server on %s", ctx.Conf().API.RestAddr)
	if err := engine.Run(ctx.Conf().API.RestAddr); err != nil {
		log.Fatalf("Error while starting API server: %s", err.Error())
	}
}
