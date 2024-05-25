package config

import "os"

type jwtOption struct {
	JWTToken string
}

type frontend struct {
	Host string
	Port string
}

type Config struct {
	Port     string
	Frontend frontend
	JWT      jwtOption
}

func Get() *Config {
	return &Config{
		Port: os.Getenv("PORT"),
		Frontend: frontend{
			Host: os.Getenv("FRONT_HOST"),
			Port: os.Getenv("FRONT_PORT"),
		},
		JWT: jwtOption{
			os.Getenv("JWT_SECRET"),
		},
	}
}
