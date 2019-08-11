package api

import (
	"log"
	"net"

	"github.com/enderian/directrd/terminals"
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
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

	buf := make([]byte, 2048)
	for {
		length, addr, err := listener.ReadFrom(buf)
		if err != nil {
			log.Printf("failed while reading bytes: %v", err)
			continue
		}

		event := &types.Event{}
		if err = proto.Unmarshal(buf[:length], event); err != nil {
			log.Printf("failed while parsing bytes: %v", err)
			log.Println(err)
			continue
		}

		switch event.Scope {
		case types.Event_Terminal:
			terminals.Event(event, addr)
		}
	}
}
