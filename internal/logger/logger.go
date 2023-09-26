package logger

import (
	"github.com/kardianos/service"
)

type Logger struct {
	errors chan error
	service.Logger
}

func New(svc service.Service) (l *Logger, err error) {
	l = new(Logger)
	l.errors = make(chan error)
	if l.Logger, err = svc.Logger(l.errors); err != nil {
		return
	}
	go func() {
		for {
			select {
			case err := <-l.errors:
				_ = l.Error(err)
			}
		}
	}()
	return
}

func (l *Logger) Info(v ...any) (err error) {
	return l.Logger.Info(v)
}

func (l *Logger) Warning(v ...any) (err error) {
	return l.Logger.Warning(v)
}

func (l *Logger) Error(v ...any) (err error) {
	return l.Logger.Error(v)
}

func (l *Logger) Infof(format string, v ...any) (err error) {
	return l.Logger.Infof(format, v)
}

func (l *Logger) Warningf(format string, v ...any) (err error) {
	return l.Logger.Warningf(format, v)
}

func (l *Logger) Errorf(format string, v ...any) (err error) {
	return l.Logger.Errorf(format, v)
}
