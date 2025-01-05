package cli

import (
	"fmt"
	"tmdb-cli/internal/api"
	"tmdb-cli/pkg/formatter"
)

// App encapsulates the CLI application with a TMDB API client.
type App struct {
	client *api.Client
}

// NewApp initializes a new App with the given TMDB API client.
func NewApp(client *api.Client) *App {
	if client == nil {
		return nil
	}
	return &App{
		client: client,
	}
}

// Execute fetches and prints movies of the specified type.
func (a *App) Execute(movieType string) error {
	if a.client == nil {
		return fmt.Errorf("client is not initialized")
	}

	movies, err := a.client.GetMovies(movieType)
	if err != nil {
		return fmt.Errorf("failed to get movies: %w", err)
	}

	formatter.PrintMovies(movies.Results)
	return nil
}
