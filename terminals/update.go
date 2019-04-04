package terminals

import (
	"fmt"
	"github.com/enderian/directrd/types"
	"log"
	"time"
)

func Update(event *types.Event) {
	err := ctx.Redis().Set(fmt.Sprintf(redisTerminalKeyAlive, event.Terminal), true, time.Second*5).Err()
	if err != nil {
		log.Println("error while updating alive for %s: %v", event.Terminal, err)
	}
}
