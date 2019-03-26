package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func status(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"name": "directrd REST API",
	})
}
