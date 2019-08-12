package api

import (
	"net/http"

	"github.com/enderian/directrd/types"
	"github.com/labstack/echo"
)

func terminalsGroup(g *echo.Group) {
	g.GET("/all", terminalsAll)
	g.GET("/:terminal", getSingleTerminal)
}

func terminalsAll(c echo.Context) error {
	var terminals []*types.Terminal
	if err := ctx.DB().Find(&terminals).Error; err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, terminals)
}

func getSingleTerminal(c echo.Context) error {
	var term []*types.Terminal
	terminal := c.Param("terminal")

	if err := ctx.DB().Where("name = ?", terminal).Find(&term).Error; err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, term)
}
