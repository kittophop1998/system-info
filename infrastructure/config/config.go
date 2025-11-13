package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type ServerConfig struct {
	Port     int    `yaml:"port"`
	GinMode  string `yaml:"gin_mode"`
	LogLevel string `yaml:"log_level"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig(env string) (*Config, error) {
	var loadErr error
	once.Do(func() {
		filename := fmt.Sprintf("configs/%s.yaml", env)
		data, err := os.ReadFile(filename)
		if err != nil {
			loadErr = fmt.Errorf("failed to read config file: %w", err)
			return
		}

		var config Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			loadErr = fmt.Errorf("failed to parse config: %w", err)
			return
		}

		cfg = &config
	})

	return cfg, loadErr
}
