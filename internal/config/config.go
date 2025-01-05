package config

import (
	"os"
)

// Config holds the configuration for the application.
type Config struct {
	BearerToken string
}

// NewConfig initializes a new Config with the given bearer token.
func NewConfig(bearerToken string) *Config {
	return &Config{
		BearerToken: bearerToken,
	}
}

// Load Loads configuration from environment variables.
func Load() (*Config, error) {
	token := os.Getenv("TMDB_BEARER_TOKEN")
	if token == "" {
		token = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MzhkYzg2YmE1ZTBkZWU0OTEwNGVkOTVkMDdjMDUyNSIsIm5iZiI6MTczMzIzNjk4Mi44NTY5OTk5LCJzdWIiOiI2NzRmMThmNjY3OWVjMTUzMjg4MjZiYmEiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.KCzSUU8Infr8oylV31eMG14F9U8ZcSC2M0ApSh_02PE"
	}

	return NewConfig(token), nil
}
