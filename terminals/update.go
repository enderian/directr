package terminals

import (
	"github.com/enderian/directrd/types"
	"log"
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
		
	case types.Event_Goodbye:
		
	}
}
