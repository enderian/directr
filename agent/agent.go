package agent

import (
	"fmt"
	"os"

	"github.com/enderian/directrd/types"
	"github.com/kardianos/service"
)

var svcConfig = &service.Config{
	Name:        "directrd_agent",
	DisplayName: "directrd agent",
	Description: "The agent process for directrd.",
}

var hostname string
var config *types.Configuration
var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	if hstn, err := os.Hostname(); err == nil {
		hostname = hstn
	} else {
		return fmt.Errorf("unable to retrieve machines hostname: %v", err)
	}

	go p.run()
	return nil
}

func (p *program) run() {
	logger.Infof("directrd agent starting")

	greet()
	go runKeepAlive()
	go runCommandReceiver()
}

func (p *program) Stop(s service.Service) error {
	logger.Infof("directrd agent stopping")

	goingDown()
	return nil
}
