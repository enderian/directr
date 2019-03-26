package api

import (
	"github.com/enderian/directrd/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var terminalUpgrader = websocket.Upgrader{}

func terminalWebSocket(c *gin.Context) {
	conn, err := terminalUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	terminal := &types.Terminal{}
	err = conn.ReadJSON(terminal)
	if err != nil {
		log.Printf("Error while reading terminal information: %s", err)
		return
	}

	log.Println(terminal)

}
