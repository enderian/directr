package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name": "directrd REST API",
	})
}
