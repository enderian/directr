package agent

import (
	"fmt"
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

func (agent *Agent) runCommandReceiver() {
	addr, _ := net.ResolveUDPAddr("udp", agent.config.API.ServiceAddr)
	addr.Port = addr.Port + 1

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panicf("failed to initialize command listener: %v", err)
	}

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
