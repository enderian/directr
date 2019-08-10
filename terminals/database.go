package terminals

import (
	"github.com/enderian/directrd/types"
	"log"
)

const redisTerminalKey = "ender.directrd.terminal.%s"
const redisTerminalKeyAlive = "ender.directrd.terminal.%s.alive"

func loadTerminals() {
	var terminals []*types.Terminal
	if err := ctx.DB().Find(&terminals).Error; err != nil {
		log.Fatalf("failed to load terminals: %v", err)
		return
	}
	log.Printf("loaded %d terminals from the database.", len(terminals))
}