package config

import (
	"encoding/json"
	"os"
)

const cfgFileName = ".gatorconfig.json"

// Config probably does something
type Config struct {
	dbURL           string
	currentUserName string
}

// Read passes cfg (struct) from homedir file
func Read() (Config, error) {
	var cfg Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}
	fileContent, err := os.ReadFile(homeDir + cfgFileName)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(fileContent, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

// SetUser -
func (c Config) SetUser() string {
	return ""
}
