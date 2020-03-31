package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	PathPrefix string
	APIPrefix  string

	ImagePath string
	ImageURL  string

	Address string
	DBPath  string
}

func readConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
