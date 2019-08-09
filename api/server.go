package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

func startApiServer() {

	if ctx.Conf().API.Disabled {
		return
	}

	engine := echo.New()
	engine.Use(middleware.Recover())

	engine.GET("/status", status)

	api := engine.Group("/api")
	usersGroup(api.Group("/users"))
	terminalsGroup(api.Group("/terminal"))

	log.Printf("Starting API server on %s", ctx.Conf().API.RestAddr)
	if err := engine.Start(ctx.Conf().API.RestAddr); err != nil {
		log.Fatalf("Error while starting API server: %s", err.Error())
	}
}
