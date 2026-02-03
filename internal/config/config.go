package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	jsonFile, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(jsonFile, &config)
	return config, err
}

func (cfg Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(cfg)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonConfig, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	os.WriteFile(fullPath, jsonConfig, 0644)
	return nil
}
