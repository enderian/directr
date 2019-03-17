package service

import "github.com/kardianos/service"

type AgentService struct {
	service service.Service
}

func (AgentService) Run() error {
	panic("implement me")
}

func (s *AgentService) Start() error {
	// Start should not block. Do the actual work async.
	go s.run()
	return nil
}
func (s *AgentService) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}
func (s *AgentService) run() {
	// Do work here
}
