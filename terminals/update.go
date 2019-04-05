package terminals

import (
	"fmt"
	"github.com/enderian/directrd/types"
	"log"
	"time"
)

func Update(event *types.Event) {
	log.Printf("updated: %v, %v", event.Terminal, event.Type)

	switch event.Type {
	case types.Event_KeepAlive:
		err := ctx.Redis().Set(fmt.Sprintf(redisTerminalKeyAlive, event.Terminal), true, time.Second*5).Err()
		if err != nil {
			log.Printf("error while updating alive for %s: %v", event.Terminal, err)
			return
		}
	case types.Event_Goodbye:
		if err := ctx.Redis().Del(fmt.Sprintf(redisTerminalKeyAlive, event.Terminal)).Err(); err != nil {
			log.Printf("error while deleting alive for %s: %v", event.Terminal, err)
			return
		}
	}
}
