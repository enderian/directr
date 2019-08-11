package agent

import (
	"os"
	"os/exec"

	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
)

func runCommandReceiver() {
	conn := incomingUDP()
	byt := make([]byte, 2048)

	for {
		length, err := conn.Read(byt)
		if err != nil {
			logger.Errorf("error while reading command: %v", err)
			return
		}

		incoming := &types.Command{}
		err = proto.Unmarshal(byt[:length], incoming)
		if err != nil {
			logger.Errorf("error while reading command: %v", err)
		}

		cmd := exec.Command(incoming.GetCommand(), incoming.GetArguments()...)
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			logger.Errorf("error while executing command: %v", err)
		}
	}
}
