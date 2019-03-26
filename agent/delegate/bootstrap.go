package main

import (
	"github.com/enderian/directrd/agent"
	"github.com/kardianos/service"
	"log"
)

func ServiceConfig() *service.Config {
	return &service.Config{
		Name:        "directrd Agent Service",
		DisplayName: "directrd Agent Service",
	}
}

func main() {
	svcConfig := ServiceConfig()

	prg := &agent.Agent{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	prg.Logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		_ = prg.Logger.Error(err)
	}
}
