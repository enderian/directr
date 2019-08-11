package agent

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/enderian/directrd/types"
	"github.com/kardianos/service"
)

func Start(c *cli.Context) error {

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return err
	}
	logger, err = s.Logger(nil)
	if err != nil {
		return err
	}

	if conf, err := types.LoadConfiguration(c); err == nil {
		config = conf
	} else {
		return fmt.Errorf("failed to load configuration: %v", err)
	}

	if c.BoolT("deamon") {
		return s.Start()
	} else {
		log.Println("starting the agent in no-deamon mode")
		return s.Run()
	}
}

func Stop(c *cli.Context) error {

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return err
	}

	logger, err = s.Logger(nil)
	if err != nil {
		return err
	}

	return s.Stop()
}
