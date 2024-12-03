package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	APIKey         string `yaml:"api_key"`
	LogLevel       string `yaml:"log_level"`
	CacheDirectory string `yaml:"cache_directory"`
	Timeout        int    `yaml:"timeout_seconds"`
	Metrics        struct {
		Enabled bool   `yaml:"enabled"`
		Port    int    `yaml:"port"`
		Path    string `yaml:"path"`
	} `yaml:"metrics"`
}

func LoadConfig() (*Config, error) {
	// Priority: Environment Variables > Config File > Default Values
	config := &Config{
		LogLevel:       "info",
		CacheDirectory: "/tmp/tmdb-cache",
		Timeout:        30,
		Metrics: struct {
			Enabled bool   `yaml:"enabled"`
			Port    int    `yaml:"port"`
			Path    string `yaml:"path"`
		}{Enabled: true, Port: 9090, Path: "/metrics"},
	}

	// Check environment variables first
	if apiKey := os.Getenv("TMDB_API_KEY"); apiKey != "" {
		config.APIKey = apiKey
	}

	// Try to load from config file
	configPath := os.Getenv("TMDB_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yaml"
	}

	configData, err := os.ReadFile(configPath)
	if err == nil {
		if err := yaml.Unmarshal(configData, config); err != nil {
			return nil, fmt.Errorf("error parsing config file: %v", err)
		}
	}

	// Validate configuration
	if config.APIKey == "" {
		return nil, fmt.Errorf("TMDB API key is required")
	}

	return config, nil
}

func (c *Config) GetCachePath() string {
	return c.CacheDirectory
}

func (c *Config) IsMetricsEnabled() bool {
	return c.Metrics.Enabled
}
