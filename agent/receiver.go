package agent

import (
	"fmt"
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
)

func (agent *Agent) runCommandReceiver() {
	conn := agent.incomingUDP()
	byt := make([]byte, 2048)

	for {
		length, err := conn.Read(byt)
		if err != nil {
			_ = agent.logger.Error(err)
			return
		}

		cmd := &types.Command{}
		err = proto.Unmarshal(byt[:length], cmd)
		if err != nil {
			_ = agent.logger.Error(err)
		}
		if cmd.Terminal != agent.hostname {
			continue
		}

		fmt.Println(cmd.Command)
	}
}
