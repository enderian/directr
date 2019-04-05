package terminals

import "github.com/enderian/directrd/types"

var ctx types.Context
var commandQueue chan types.Command

func Setup(context types.Context) {
	ctx = context
	go setupRedisSubscriber()
}

func SetupQueue(queue chan types.Command) {
	commandQueue = queue
}
