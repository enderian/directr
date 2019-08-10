package terminals

import (
	"fmt"
	"github.com/enderian/directrd/types"
	"log"
	"time"
)

func Update(event *types.Event) {
	log.Printf("updated: %v, %v", event.Terminal, event.Type)

	var terminal = &types.Terminal{}
	err := ctx.DB().Where("name = ?", event.Terminal).Find(terminal)
	if err != nil {
		log.Printf("terminal attempted to connect but does not exist: %v: %v", event.Terminal, err)
		return
	}

	switch event.Type {
	case types.Event_KeepAlive:
		err := ctx.Redis().Set(fmt.Sprintf(redisTerminalKeyAlive, terminal.Hostname), true, time.Second*5).Err()
		if err != nil {
			log.Printf("error while updating alive for %s: %v", event.Terminal, err)
			return
		}
	case types.Event_Goodbye:
		if err := ctx.Redis().Del(fmt.Sprintf(redisTerminalKeyAlive, terminal.Hostname)).Err(); err != nil {
			log.Printf("error while deleting alive for %s: %v", event.Terminal, err)
			return
		}
	}
}
