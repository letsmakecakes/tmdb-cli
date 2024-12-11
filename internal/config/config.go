package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

var (
	tmdbAPIKey string
)

// LoadConfig Loads environment variables from .env file
func LoadConfig() error {
	// Load .env file in the current directory, if it exists
	if err := godotenv.Load(); err != nil {
		return errors.New("Error loading config .env file " + err.Error())
	}

	// Retrieve TMDB API Key
	tmdbAPIKey = os.Getenv("TMDB_API_KEY")
	if tmdbAPIKey == "" {
		return errors.New("TMDB_API_KEY is not set in the environment")
	}

	return nil
}
