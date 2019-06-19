package agent

import (
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
	"time"
)

func (agent *Agent) runKeepAlive() {
	conn := agent.outgoingUDP()

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

	event := &types.Event{
		Terminal: agent.hostname,
		Scope:    types.Event_Terminal,
		Type:     types.Event_Goodbye,
	}

	msg, err := proto.Marshal(event)
	if err != nil {
		panic(err)
	}
	conn := agent.outgoingUDP()
	if _, err = conn.Write(msg); err != nil {
		_ = agent.logger.Errorf("failed on sending goodbye: %v", err)
	}

}
