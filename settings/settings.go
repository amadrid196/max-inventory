package settings

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed settings.yaml
var settingsFile []byte

// Data structure to hold the settings data
type DataBaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// Settings struct to hold the settings
type Settings struct {
	Host  string         `yaml:"host"`
	Port  int            `yaml:"port"`
	Debug bool           `yaml:"debug"`
	DB    DataBaseConfig `yaml:"database"`
}

func New() (*Settings, error) {
	var s Settings

	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
