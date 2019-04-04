package agent

import (
	"github.com/codegangsta/cli"
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
	"github.com/kardianos/service"
	"log"
	"net"
)

type Agent struct {
	Logger service.Logger
}

func Setup(_ *cli.Context) error {
	svcConfig := &service.Config{
		Name:        "directrd agent service",
		DisplayName: "directrd agent service",
	}

	prg := &Agent{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	prg.Logger, err = s.Logger(nil)
	if err != nil {
		return err
	}
	return s.Run()
}

func (agent *Agent) Start(s service.Service) error {
	go agent.run()
	return nil
}

func (agent *Agent) Stop(s service.Service) error {
	return nil
}

func (agent *Agent) run() {
	addr := &net.UDPAddr{
		Port: 12056,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panicf("failed to initialize internal listener: %v", err)
	}

	event := &types.Event{
		Terminal: "cslab-12",
		Scope:    types.Event_Terminal,
		Type:     types.Event_KeepAlive,
	}

	msg, err := proto.Marshal(event)
	if err != nil {
		_ = agent.Logger.Error(err)
	}
	if _, err = conn.Write(msg); err != nil {
		_ = agent.Logger.Error(err)
	}
}
