package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

func ReadConfig() (Config, error) {
	cfgLocation, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(cfgLocation)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	cfgLocation := filepath.Join(home, configFileName)
	return cfgLocation, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	cfgLocation, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(cfgLocation)
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
