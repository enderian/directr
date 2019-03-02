package api

import (
	"github.com/kataras/iris"
	recoverMiddleware "github.com/kataras/iris/middleware/recover"
	"log"
)

func startApiServer() {

	if ctx.Conf().API.Disabled {
		return
	}

	irs := iris.New()
	irs.Use(recoverMiddleware.New())
	irs.Logger().SetOutput(ctx.Logger())

	//TODO Define all the routes
	api := irs.Party("/api")
	api.Get("/", status)

	log.Printf("Starting API server on %s", ":8080")
	if err := irs.Run(iris.Addr(":8080")); err != nil {
		log.Fatalf("Error while starting API server: %s", err.Error())
	}
}
