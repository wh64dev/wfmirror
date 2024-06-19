package config

import (
	"encoding/json"
	"os"
)

type service struct {
	Name string
}

type jwtOption struct {
	PrivKey string
}

type GlobalConf struct {
	DataDir string `json:"data_dir"`
}

type DockerConf struct {
	PreUsername string
	PrePassword string
}

type Config struct {
	Port        string
	AllowOrigin string
	Service     service
	JWT         jwtOption
	Global      GlobalConf
	Docker      DockerConf
}

func Get() *Config {
	file, err := os.ReadFile("./temp/config.json")
	if err != nil {
		file = nil
	}

	var data GlobalConf
	err = json.Unmarshal(file, &data)
	if err != nil {
		data = GlobalConf{}
	}

	return &Config{
		Port:        os.Getenv("PORT"),
		AllowOrigin: os.Getenv("ALLOW_ORIGIN"),
		Service: service{
			Name: os.Getenv("SERVICE_NAME"),
		},
		JWT: jwtOption{
			PrivKey: os.Getenv("JWT_SECRET"),
		},
		Global: data,
		Docker: DockerConf{
			PreUsername: os.Getenv("PRE_USERNAME"),
			PrePassword: os.Getenv("PRE_PASSWORD"),
		},
	}
}

func Set(confType string, value string) error {
	file, err := os.ReadFile("./temp/config.json")
	if err != nil {
		return err
	}

	var data GlobalConf
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	switch confType {
	case "dir":
		data.DataDir = value
	}

	bytes, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	err = os.WriteFile("./temp/config.json", bytes, 0755)
	if err != nil {
		return err
	}

	return nil
}

func LoadDefault() ([]byte, error) {
	var data = GlobalConf{
		DataDir: "data",
	}

	bytes, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
