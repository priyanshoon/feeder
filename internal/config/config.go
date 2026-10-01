package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(userPath, configFileName)

	return path, nil
}

func write(cfg Config) error {
	fullpath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullpath)
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
