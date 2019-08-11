package api

import (
	"net/http"

	"github.com/enderian/directrd/delegation"
	"github.com/enderian/directrd/types"
	"github.com/labstack/echo"
)

func usersGroup(g *echo.Group) {
	g.GET("/:username", getSingleUser)
}

func getSingleUser(c echo.Context) error {
	username := c.Param("username")

	user := &types.User{Username: username}

	_ = delegation.FillLdap(user)
	// Hide DN
	user.DN = ""

	return c.JSON(http.StatusOK, user)
}
