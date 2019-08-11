package agent

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
)

func greet() {
	conn := outgoingUDP()

	event := &types.Event{
		Terminal: hostname,
		Scope:    types.Event_Terminal,
		Type:     types.Event_Greetings,
		Data: map[string]string{
			"os": runtime.GOOS,
		},
	}

	msg, err := proto.Marshal(event)
	if err != nil {
		log.Fatalf("failed on marshaling greeting: %v", err)
	}
	if _, err = conn.Write(msg); err != nil {
		log.Fatalf("failed on sending greeting: %v", err)
	}
}

func runKeepAlive() {
	conn := outgoingUDP()

	for {
		event := &types.Event{
			Terminal: hostname,
			Scope:    types.Event_Terminal,
			Type:     types.Event_KeepAlive,
		}

		msg, err := proto.Marshal(event)
		if err != nil {
			log.Fatalf("failed on marshaling keep_alive: %v", err)
		}
		if _, err = conn.Write(msg); err != nil {
			log.Printf("failed on sending keep_alive: %v", err)
			continue
		}

		time.Sleep(time.Second * 2)
	}
}

func goingDown() error {

	event := &types.Event{
		Terminal: hostname,
		Scope:    types.Event_Terminal,
		Type:     types.Event_Goodbye,
	}

	msg, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	conn := outgoingUDP()
	if _, err = conn.Write(msg); err != nil {
		return fmt.Errorf("failed on sending goodbye: %v", err)
	}
	return nil
}
