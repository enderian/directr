package api

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func startApiServer() {

	if ctx.Conf().API.Disabled {
		return
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	e.GET("/status", status)

	api := e.Group("/api")
	usersGroup(api.Group("/users"))
	terminalsGroup(api.Group("/terminals"))

	log.Printf("Starting API server on %s", ctx.Conf().API.RestAddr)
	if err := e.Start(ctx.Conf().API.RestAddr); err != nil {
		log.Fatalf("Error while starting API server: %s", err.Error())
	}
}
