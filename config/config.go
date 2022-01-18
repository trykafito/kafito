package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Port      string `json:"port"`
		SecretKey string `json:"secret_key"`
		Mongo     Mongo  `json:"mongo"`
	}

	Mongo struct {
		Host     string `json:"host"`
		DB       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
)

func Read(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := new(Config)
	if err := json.Unmarshal(data, c); err != nil {
		return nil, err
	}

	return c, nil
}
