package main

import (
	"github.com/enderian/directrd/agent"
	"github.com/kardianos/service"
	"log"
)

func main() {
	program := &agent.Program{}
	srv, err := service.New(program, agent.ServiceConfig())
	if err != nil {
		log.Printf("Unable to initialize service: %s", err.Error())
	}
	if err := srv.Run(); err != nil {
		log.Printf("Unable to run service: %s", err.Error())
	}
}
