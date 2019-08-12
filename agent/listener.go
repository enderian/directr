package agent

import (
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

func changeListener() {
	conn := outgoingUDP()
	var sessionID uintptr = 0xFFFFFFFF

	for {
		newSessionID := getSessionID()

		if newSessionID != sessionID {
			sessionID = newSessionID
			if sessionID == 0xFFFFFFFF {
				event := &types.Event{
					Terminal: hostname,
					Scope:    types.Event_Terminal,
					Type:     types.Event_SessionEnd,
				}

				msg, err := proto.Marshal(event)
				if err != nil {
					log.Fatalf("failed on marshaling event session_end: %v", err)
				}
				if _, err = conn.Write(msg); err != nil {
					log.Printf("failed on sending event session_end: %v", err)
					continue
				}
			} else {
				ret, username := getUsername(sessionID)

				if ret != 0 {
					event := &types.Event{
						Terminal: hostname,
						Scope:    types.Event_Terminal,
						Type:     types.Event_SessionStart,
						Data: map[string]string{
							"username": username,
						},
					}

					msg, err := proto.Marshal(event)
					if err != nil {
						log.Fatalf("failed on marshaling event session_start: %v", err)
					}
					if _, err = conn.Write(msg); err != nil {
						log.Printf("failed on sending event session_start: %v", err)
						continue
					}

				} else {
					log.Fatal("Error while retrieving username!")
				}
			}
		}
		time.Sleep(time.Second)
	}
}
