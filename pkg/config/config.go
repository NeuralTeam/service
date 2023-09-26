package config

import "github.com/kardianos/service"

type Config service.Config

func New() (cfg *Config) {
	cfg = new(Config)
	return
}
