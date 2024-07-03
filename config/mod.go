package config

import (
	"encoding/json"
	"os"
)

type service struct {
	Name string
}

type GlobalConf struct {
	DataDir string
}

type External struct {
	Version string `json:"version"`
}

type Config struct {
	Port        string
	AllowOrigin string
	Service     service
	Dist        External
	Global      GlobalConf
}

func Get() *Config {
	file, err := os.ReadFile("./package.json")
	if err != nil {
		file = nil
	}

	var data External
	err = json.Unmarshal(file, &data)
	if err != nil {
		data = External{}
	}

	return &Config{
		Port: os.Getenv("PORT"),
		Dist: data,
		Service: service{
			Name: os.Getenv("SERVICE_NAME"),
		},
		Global: GlobalConf{
			DataDir: os.Getenv("DATA_DIR"),
		},
	}
}
