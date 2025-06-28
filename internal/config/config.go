package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const cfgFileName = ".gatorconfig.json"

// Config probably does something
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read also does something
func Read() (Config, error) {
	var cfg Config
	path, err := getConfigPath()
	if err != nil {
		return cfg, err
	}
	fmt.Println("Read call on:", path)
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(fileContent, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// SetUser couldn't possibly set the user??
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	path, _ := getConfigPath()
	fileContent, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, fileContent, 0777)
	if err != nil {
		return err
	}
	return nil
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + cfgFileName, nil
}
