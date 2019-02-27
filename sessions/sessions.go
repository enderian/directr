package sessions

import (
	"context"
	"github.com/enderian/directrd/types"
)

var ctx context.Context
var conf *types.Configuration

func Setup(context context.Context, configuration *types.Configuration) {
	conf = configuration
	ctx = context
}
