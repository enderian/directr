package terminals

import (
	"encoding/json"
	"fmt"
	"github.com/enderian/directrd/types"
	"log"
)

const redisTerminalKey = "ender.directrd.terminal.%s"
const redisTerminalKeyAlive = "ender.directrd.terminal.%s.alive"

func loadTerminals() {
	var terminals []*types.Terminal
	if err := ctx.DB().Model(&terminals).Select(); err != nil {
		log.Fatalf("failed to load terminals: %v", err)
		return
	}
	for _, term := range terminals {
		if err := updateTerminal(term); err != nil {
			log.Fatalf("failed to load terminal %s: %v", term.Name, err)
		}
	}
	log.Printf("loaded %d terminals from the database.", len(terminals))
}

func findTerminal(hostname string) (*types.Terminal, error) {
	res := ctx.Redis().Get(fmt.Sprintf(redisTerminalKey, hostname))
	if res.Err() != nil {
		return nil, res.Err()
	}

	byt, err := res.Bytes()
	if err != nil {
		return nil, err
	}

	terminal := &types.Terminal{}
	err = json.Unmarshal(byt, terminal)
	return terminal, err
}

func updateTerminal(terminal *types.Terminal) error {
	js, _ := json.Marshal(terminal)
	return ctx.Redis().Set(fmt.Sprintf(redisTerminalKey, terminal.Hostname), js, 0).Err()
}

func insertTerminal(terminal *types.Terminal) error {
	_, err := ctx.DB().Model(terminal).Insert()
	return err
}
