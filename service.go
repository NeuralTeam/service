package service

import (
	"github.com/NeuralTeam/service/internal/service"
	"github.com/NeuralTeam/service/pkg/config"
	"github.com/NeuralTeam/service/pkg/events"
	"github.com/NeuralTeam/service/pkg/status"
)

type Service interface {
	Install() (err error)
	Uninstall() (err error)
	Reinstall() (err error)

	Run() (err error)
	Restart() (err error)

	String() (name string)
	Status() (status status.Status, err error)
	Platform() (platform string)

	Info(v ...any) (err error)
	Warning(v ...any) (err error)
	Error(v ...any) (err error)

	Infof(format string, v ...any) (err error)
	Warningf(format string, v ...any) (err error)
	Errorf(format string, v ...any) (err error)
}

func New(cfg *config.Config, events *events.Events) (svc Service, err error) {
	return service.New(cfg, events)
}
