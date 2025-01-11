package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

// Path to the configuration file.
var configPath = filepath.Join(xdg.ConfigHome, "powermenu", "config.yaml")

// Resolve returns configuration from file if it exists and is valid.
// Otherwise, default configuration is returned.
func Resolve() Config {
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return defaultConfig
	}

	var config Config

	if yaml.Unmarshal(configBytes, &config) != nil {
		return defaultConfig
	}

	return config
}
