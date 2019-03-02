package api

import (
	"github.com/enderian/directrd/types"
)

var ctx types.Context
var event = make(chan types.Event)

func Setup(context types.Context) {
	ctx = context
	go startApiServer()
}
