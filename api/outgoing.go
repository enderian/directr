package api

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/enderian/directrd/terminals"
	"github.com/enderian/directrd/types"
	"github.com/golang/protobuf/proto"
)

var commandQueue = make(chan types.Command)

func outgoingUDP(destination string) *net.UDPConn {
	_, portStr, err := net.SplitHostPort(ctx.Conf().API.ServiceAddr)
	if err != nil {
		log.Panicf("failed to split address: %v", err)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Panicf("failed to parse port %s: %v", portStr, err)
	}

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%v", destination, port+1))
	if err != nil {
		log.Panicf("failed to parse server address: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panicf("failed to initialize sender: %v", err)
	}
	return conn
}

func startInternalOutgoing() {
	terminals.SetupQueue(commandQueue)

	for {
		cmd := <-commandQueue
		addr, err := terminals.GetAddr(cmd.GetTerminal())
		if err != nil {
			log.Printf("error while sending command: %v", err)
			continue
		}

		conn := outgoingUDP(addr)
		msg, err := proto.Marshal(&cmd)
		if err != nil {
			log.Printf("error while sending command: %v", err)
			continue
		}
		conn.Write(msg)

		log.Printf("sent command to %s", addr)
	}
}
