package agent

import (
	"github.com/codegangsta/cli"
	"github.com/enderian/directrd/types"
	"github.com/kardianos/service"
	"log"
	"os"
)

type Agent struct {
	hostname string
	config   *types.Configuration
	logger   service.Logger
}

func Setup(c *cli.Context) error {
	svcConfig := &service.Config{
		Name:        "directrd agent service",
		DisplayName: "directrd agent service",
	}

	prg := &Agent{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	prg.logger, err = s.Logger(nil)
	if err != nil {
		return err
	}

	if conf, err := types.LoadConfiguration(c); err == nil {
		prg.config = conf
	} else {
		_ = prg.logger.Errorf("failed to load configuration: $v", err)
		return err
	}

	return s.Run()
}

func (agent *Agent) Start(s service.Service) error {
	_ = agent.logger.Info("directrd agent starting")

	go agent.run()

	return nil
}

func (agent *Agent) Stop(s service.Service) error {
	_ = agent.logger.Info("directrd agent stopping")
	agent.goingDown()
	return nil
}

func (agent *Agent) run() {
	if hostname, err := os.Hostname(); err == nil {
		agent.hostname = hostname
	} else {
		_ = agent.logger.Errorf("unable to retrieve machines hostname: $v", err)
		return
	}

	go agent.runKeepAlive()
}
