package config

import (
	"os"
)

type service struct {
	Name string
}

type jwtOption struct {
	PrivKey string
}

type Config struct {
	Port        string
	AllowOrigin string
	Service     service
	JWT         jwtOption
}

func Get() *Config {
	return &Config{
		Port:        os.Getenv("PORT"),
		AllowOrigin: os.Getenv("ALLOW_ORIGIN"),
		Service: service{
			Name: os.Getenv("SERVICE_NAME"),
		},
		JWT: jwtOption{
			PrivKey: os.Getenv("JWT_SECRET"),
		},
	}
}
