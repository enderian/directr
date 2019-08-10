package terminals

import (
	"log"

	"github.com/enderian/directrd/types"
)

func Update(event *types.Event) {

	var terminal = &types.Terminal{}
	if err := ctx.DB().Where("hostname = ?", event.Terminal).Find(terminal).Error; err != nil {
		log.Printf("terminal attempted to connect but does not exist: %v: %v", event.Terminal, err)
		return
	}

	switch event.Type {
	case types.Event_KeepAlive:
		if terminal.Status == types.StatusOffline {
			terminal.Status = types.StatusOnline
			terminal.SaveStatus()
		}
	case types.Event_Goodbye:
		terminal.Status = types.StatusOffline
		ctx.DB().Update(terminal)
	}
}
