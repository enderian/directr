package directoryRadius

import (
	"context"
	"ender.gr/directrd/types"
)

var ctx context.Context
var conf *types.Configuration

func Setup(context context.Context, configuration *types.Configuration) {
	conf = configuration
	ctx = context
	go startAuthServer()
	go startAccServer()
}
