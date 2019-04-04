package api

import (
	"bufio"
	"github.com/enderian/directrd/terminals"
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

func startInternal() {
	addr, err := net.ResolveUDPAddr("udp", ctx.Conf().API.ServiceAddr)
	if err != nil {
		log.Panicf("failed to parse service addr: %v", err)
	}

	log.Printf("Starting internal service server on %v", ctx.Conf().API.ServiceAddr)
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panicf("failed to initialize internal listener: %v", err)
	}

	writer := bufio.NewReader(listener)
	byt := make([]byte, 1024)
	for {
		num, err := writer.Read(byt)
		if err != nil {
			return
		}
		event := &types.Event{}
		if err = proto.Unmarshal(byt[:num], event); err != nil {
			log.Println(err)
			continue
		}

		switch event.Scope {
		case types.Event_Terminal:
			terminals.Update(event)
		}
	}
}

func startInternalOutgoing() {
	terminals.SetupQueue(commandQueue)

	for {
		cmd := <-commandQueue
		_ = cmd //FIXME Send to corresponding terminal!
	}
}
