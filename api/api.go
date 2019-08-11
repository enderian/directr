package api

import (
	"github.com/enderian/directrd/types"
)

var ctx types.Context

func Setup(context types.Context) {
	ctx = context

	go startApiServer()
	go startInternal()
	go startInternalOutgoing()
}
