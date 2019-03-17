package agent

import "github.com/kardianos/service"

func ServiceConfig() *service.Config {
	return &service.Config{
		Name:        "directrd Agent Service",
		DisplayName: "directrd Agent Service",
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	return nil
}

func (p *Program) run() {

}
