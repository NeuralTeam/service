package service

import (
	"github.com/NeuralTeam/service/pkg/status"
	"github.com/kardianos/service"
)

func (svc *Service) String() (name string) {
	name = svc.Service.String()
	return
}

func (svc *Service) Status() (status status.Status, err error) {
	var s service.Status
	s, err = svc.Service.Status()
	status = status.Status(s)
	return
}

func (svc *Service) Platform() (platform string) {
	platform = svc.Service.Platform()
	return
}
