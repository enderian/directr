package radius

import (
	"github.com/enderian/directrd/types"
)

var ctx types.Context

func Setup(context types.Context) {
	ctx = context
	if ctx.Conf() == nil || ctx.Conf().Radius.SharedSecret == "" {
		return
	}
	go startAuthServer()
	go startAccServer()
}
