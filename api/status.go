package api

import "github.com/kataras/iris"

func status(context iris.Context) {
	_, _ = context.JSON(map[string]interface{}{
		"name": "directrd REST API",
	})
}
