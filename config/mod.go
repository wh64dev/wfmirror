package config

import (
	"os"
	"strconv"
)

type jwtOption struct {
	JWTToken string
}

type frontend struct {
	Host  string
	Port  string
	Title string
}

type Config struct {
	Port        string
	AllowOrigin string
	HashCount   int
	Frontend    frontend
	JWT         jwtOption
}

func Get() *Config {
	hc, _ := strconv.ParseInt(os.Getenv("HASH_COUNT"), 10, 32)

	return &Config{
		Port: os.Getenv("PORT"),
		Frontend: frontend{
			Host:  os.Getenv("FRONT_HOST"),
			Port:  os.Getenv("FRONT_PORT"),
			Title: os.Getenv("FRONT_TITLE"),
		},
		AllowOrigin: os.Getenv("ALLOW_ORIGIN"),
		HashCount:   int(hc),
		JWT: jwtOption{
			os.Getenv("JWT_SECRET"),
		},
	}
}
