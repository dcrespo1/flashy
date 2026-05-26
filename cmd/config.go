package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DecksDir string `json:"decks_dir"`
}

func flashyDir() (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(homeDir, ".flashy"), nil
}

func configPath() (string, error) {
    base, err := flashyDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(base, "config.json"), nil
}

func defaultDecksDir() (string, error) {
    base, err := flashyDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(base, "decks"), nil
}

func SaveConfig(cfg Config) error  {
	jsonData, err := json.MarshalIndent(cfg, "", " ")
	if err != nil{
		return err
	}

	fullPath, err := configPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(fullPath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, jsonData, 0644)
	if err != nil{
		return err
	}
	return nil
}


func LoadConfig() (Config, error) {
	fullPath, err := configPath()
	if err != nil {
    	return Config{}, err
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, err
} 


func ConfigExists() bool {
    fullPath, err := configPath()
    if err != nil {
        return false
    }
    _, err = os.Stat(fullPath)
    return err == nil
}

