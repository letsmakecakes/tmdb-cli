package main

import (
	"flag"
	"log"
	"tmdb-cli/internal/api"
	"tmdb-cli/internal/cli"
	"tmdb-cli/internal/config"
)

func main() {
	// Parse command-line flags
	movieType := parseFlags()

	// Load configuration
	cfg := loadConfig()

	// Initialize API client and CLI app
	client := api.NewClient(cfg.BearerToken)
	app := cli.NewApp(client)

	// Execute the app
	if err := app.Execute(*movieType); err != nil {
		log.Fatalf("Error executing app: %v", err)
	}
}

// parseFlags parses command-line flags and returns the movie type.
func parseFlags() *string {
	movieType := flag.String("type", "", "Type of movies to fetch (playing/popular/top/upcoming)")
	flag.Parse()

	if *movieType == "" {
		log.Fatal("Please provide a movie type using --type flag")
	}

	return movieType
}

// loadConfig loads the configuration from environment variables.
func loadConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
