package config

import "sync"

// Source represents different configuration sources
type Source string

const (
	SourceEnvFile     Source = "env_file"
	SourceEnvironment Source = "environment"
	SourceConfigFile  Source = "config_file"
)

// Config represents the application configuration
type Config struct {
	// Mutex to ensure thread-safe access
	mu sync.RWMutex

	//
}
