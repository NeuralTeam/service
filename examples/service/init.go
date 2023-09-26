package main

import (
	"github.com/NeuralTeam/service"
	"github.com/NeuralTeam/service/pkg/config"
	"github.com/NeuralTeam/service/pkg/events"
	"time"
)

var (
	Events  *events.Events
	Config  *config.Config
	Service service.Service
)

func init() {
	Config = config.New()
	Config.Name = `GoServiceExample`
	Config.DisplayName = `Go Service Example`
	Config.Description = `This is an example Go service that outputs log messages.`

	Events = events.New(
		func() {
			for t := range time.Tick(time.Second) {
				_ = Service.Infof(`still running at %v`, t)
			}
		},
		func() {
			if err := Service.Uninstall(); err != nil {
				_ = Service.Error(err)
			}
		},
	)

	var err error
	if Service, err = service.New(Config, Events); err != nil {
		panic(err)
	}
	if err = Service.Install(); err != nil {
		_ = Service.Error(err)
	}
}
