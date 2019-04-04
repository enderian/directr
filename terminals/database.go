package terminals

import (
	"encoding/json"
	"fmt"
	"github.com/enderian/directrd/types"
)

const redisTerminalKey = "ender.directrd.terminal.%s"
const redisTerminalKeyAlive = "ender.directrd.terminal.%s.alive"

func loadTerminals() error {
	return nil
}

func findTerminal(terminal *types.Terminal) error {
	res := ctx.Redis().Get(fmt.Sprintf(redisTerminalKey, terminal.Name))
	if res.Err() != nil {
		return res.Err()
	}

	byt, err := res.Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(byt, terminal)
}

func updateTerminal(terminal *types.Terminal) error {
	js, _ := json.Marshal(terminal)
	return ctx.Redis().Set(fmt.Sprintf(redisTerminalKey, terminal.Name), js, 0).Err()
}

func insertTerminal(terminal *types.Terminal) error {
	_, err := ctx.DB().Model(terminal).Insert()
	return err
}
