package config

import (
	
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Jobs []Job `yaml:"jobs"`
}

// Job represents a replication job
type Job struct {
	Name        string     `yaml:"name"`
	Source      Connection `yaml:"source"`
	Destination Connection `yaml:"destination"`
	Schedule    string     `yaml:"schedule"`
}
type Connection struct {
	Type             string `yaml:"type"`
	DSN              string `yaml:"dsn"`
	IncrementalField string `yaml:"incremental_field,omitempty"`
	Table            string `yaml:"table"`
}

func LoadFile() (*Config, error) {
	configs, err := os.ReadFile("config/Db_config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(configs, &config); err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	return &config, nil
}

func ExpandDSN(dsn string) string {
    return os.ExpandEnv(dsn)
}