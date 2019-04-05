package terminals

import (
	"github.com/enderian/directrd/database"
	"log"
)

func setupRedisSubscriber() {
	cloneCtx := database.SetupRedis(ctx)
	channel := cloneCtx.Redis().PSubscribe("__key*__:ender.directrd.*.alive").Channel()
	for {
		msg := <-channel
		switch msg.Payload {
		case "expired":
			log.Println(msg.Channel)
		}
	}
}
