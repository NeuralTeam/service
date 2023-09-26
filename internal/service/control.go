package service

import (
	"github.com/kardianos/service"
	"os"
	"strings"
)

func (svc *Service) Install() (err error) {
	if err = svc.Service.Install(); err != nil {
		if strings.Contains(err.Error(), `already exists`) {
			err = nil
		}
		return
	}
	if err = svc.Info(`installing`); err != nil {
		return
	}
	return
}

func (svc *Service) Uninstall() (err error) {
	if err = svc.Service.Uninstall(); err != nil {
		if strings.Contains(err.Error(), `not`) {
			err = nil
		}
		return
	}
	if err = svc.Info(`uninstalling`); err != nil {
		return
	}
	return
}

func (svc *Service) Reinstall() (err error) {
	if err = svc.Uninstall(); err != nil {
		return
	}
	if err = svc.Install(); err != nil {
		return
	}
	return
}

func (svc *Service) Run() (err error) {
	p := service.Platform()
	if err = svc.Infof(`running %v`, p); err != nil {
		return
	}
	if err = svc.Service.Run(); err != nil {
		return
	}
	return
}

func (svc *Service) Restart() (err error) {
	if err = svc.Info(`restarting`); err != nil {
		return
	}
	if err = svc.Service.Restart(); err != nil {
		return
	}
	return
}

func (svc *Service) Start(s service.Service) (err error) {
	if service.Interactive() {
		if err = svc.Info(`starting in terminal`); err != nil {
			return
		}
	} else {
		if err = svc.Info(`starting under service manager`); err != nil {
			return
		}
	}
	go func() {
		defer func() {
			_ = svc.Stop(svc.Service)
		}()
		svc.Events.Start()
	}()
	return
}

func (svc *Service) Stop(s service.Service) (err error) {
	if err = svc.Info(`stopping`); err != nil {
		return
	}
	defer func() {
		if service.Interactive() {
			os.Exit(0)
		}
		_ = svc.Service.Stop()
	}()
	svc.Events.Stop()
	return
}
