package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return write(*cfg)
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

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}