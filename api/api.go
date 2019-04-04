package api

import (
	"github.com/enderian/directrd/types"
)

var ctx types.Context
var commandQueue = make(chan types.Command)

func Setup(context types.Context) {
	ctx = context

	go startApiServer()
	go startInternal()
	go startInternalOutgoing()
}
