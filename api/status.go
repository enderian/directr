package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name": "directrd REST API",
	})
}
