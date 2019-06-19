package agent

import (
	"log"
	"net"
)

func (agent *Agent) incomingUDP() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", agent.config.API.ServiceAddr)
	if err != nil {
		log.Panicf("failed to parse server address: %v", err)
	}
	addr.IP = nil
	addr.Port += 1
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panicf("failed to initialize receiver: %v", err)
	}
	return conn
}

func (agent *Agent) outgoingUDP() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", agent.config.API.ServiceAddr)
	if err != nil {
		log.Panicf("failed to parse server address: %v", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panicf("failed to initialize sender: %v", err)
	}
	return conn
}
