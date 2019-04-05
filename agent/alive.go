package agent

import (
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"time"
)

func (agent *Agent) runKeepAlive() {
	addr := &net.UDPAddr{
		Port: 12056,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panicf("failed to initialize alive sender: %v", err)
	}

	for {
		event := &types.Event{
			Terminal: agent.hostname,
			Scope:    types.Event_Terminal,
			Type:     types.Event_KeepAlive,
		}

		msg, err := proto.Marshal(event)
		if err != nil {
			panic(err)
		}
		if _, err = conn.Write(msg); err != nil {
			_ = agent.logger.Errorf("failed on sending keep_alive: %v", err)
			continue
		}

		time.Sleep(time.Second * 2)
	}
}

func (agent *Agent) goingDown() {
	addr := &net.UDPAddr{
		Port: 12056,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panicf("failed to initialize goodbye sender: %v", err)
	}

	event := &types.Event{
		Terminal: agent.hostname,
		Scope:    types.Event_Terminal,
		Type:     types.Event_Goodbye,
	}

	msg, err := proto.Marshal(event)
	if err != nil {
		panic(err)
	}
	if _, err = conn.Write(msg); err != nil {
		_ = agent.logger.Errorf("failed on sending goodbye: %v", err)
	}

}
