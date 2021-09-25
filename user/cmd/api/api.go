package api

import (
	"github.com/apascualco/apascualco-user/internal/platform/server"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Host string `default:"localhost"`
	Port uint   `default:"8080"`
}

func Run() error {
	cfg, err := env()
	if err != nil {
		return err
	}

	srv := server.New(cfg.Host, cfg.Port)
	srv.ConfigureSwagger()
	srv.RegisterRoutes()
	return srv.Run()
}

func env() (config, error) {
	var cfg config
	err := envconfig.Process("USER", &cfg)
	if err != nil {
		return config{}, err
	}
	return cfg, err
}
