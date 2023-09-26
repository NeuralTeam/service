package service

import (
	"github.com/NeuralTeam/service/internal/logger"
	"github.com/NeuralTeam/service/pkg/config"
	"github.com/NeuralTeam/service/pkg/events"
	"github.com/kardianos/service"
)

type Service struct {
	*events.Events
	*logger.Logger
	service.Service
}

func New(cfg *config.Config, events *events.Events) (svc *Service, err error) {
	svc = new(Service)
	svc.Events = events
	if svc.Service, err = service.New(svc, (*service.Config)(cfg)); err != nil {
		return
	}
	if svc.Logger, err = logger.New(svc.Service); err != nil {
		return
	}
	return
}
