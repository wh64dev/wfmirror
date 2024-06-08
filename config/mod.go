package config

import (
	"os"
)

type jwtOption struct {
	PrivKey string
}

type Config struct {
	Port        string
	AllowOrigin string
	JWT         jwtOption
}

func Get() *Config {
	return &Config{
		Port:        os.Getenv("PORT"),
		AllowOrigin: os.Getenv("ALLOW_ORIGIN"),
		JWT: jwtOption{
			PrivKey: os.Getenv("JWT_SECRET"),
		},
	}
}
