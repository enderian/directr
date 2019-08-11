package terminals

import (
	"fmt"
	"log"

	"github.com/enderian/directrd/types"
)

func loadTerminals() {
	var terminals []*types.Terminal
	if err := ctx.DB().Find(&terminals).Error; err != nil {
		log.Fatalf("failed to load terminals: %v", err)
		return
	}
	log.Printf("loaded %d terminals from the database.", len(terminals))
}

const ipKey = "directrd.terminal.%s.addr"

func GetAddr(terminal string) (string, error) {
	return ctx.Redis().Get(fmt.Sprintf(ipKey, terminal)).Result()
}
