package api

import (
	"net/http"

	"github.com/enderian/directrd/types"
	"github.com/labstack/echo"
)

func terminalsGroup(g *echo.Group) {
	g.GET("/all", terminalsAll)
}

func terminalsAll(c echo.Context) error {
	var terminals []*types.Terminal
	if err := ctx.DB().Find(&terminals).Error; err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, terminals)
}
