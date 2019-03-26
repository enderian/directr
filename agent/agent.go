package agent

import (
	"github.com/enderian/directrd/types"
	"github.com/gorilla/websocket"
	"github.com/kardianos/service"
	"net/url"
	"os"
)

type Agent struct {
	Logger service.Logger
}

func (agent *Agent) Start(s service.Service) error {
	go agent.run()
	return nil
}

func (agent *Agent) Stop(s service.Service) error {
	return nil
}

func (agent *Agent) run() {

	addr := "localhost:8080"
	u := url.URL{Scheme: "ws", Host: addr, Path: "/api/terminal/ws"}
	_ = agent.Logger.Info("Connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		_ = agent.Logger.Errorf("Error while connecting to WebSocket: %s", err)
		return
	}
	defer func() {
		_ = conn.Close()
		_ = agent.Logger.Error("Service terminated.")
	}()

	hostname, err := os.Hostname()
	if err != nil {
		_ = agent.Logger.Errorf("Error while retrieving hostname: %s", err)
		return
	}
	_ = conn.WriteJSON(&types.Terminal{
		Hostname: hostname,
	})

	cmd := &types.Command{}
	for {
		err := conn.ReadJSON(cmd)
		if err != nil {
			_ = agent.Logger.Errorf("Error while parsing incoming command: %s", err)
			return
		}
		switch cmd.Type {
		case types.CommandLogout:
			logout(cmd.Arguments[0])
		case types.CommandRestart:
			restart()
		case types.CommandShutdown:
			shutdown()
		}
	}
}
