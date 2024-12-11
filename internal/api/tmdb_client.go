package api

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://api.themoviedb.org/3"
)

type TMDBClient struct {
	apikey     string
	httpClient *http.Client
}

// NewTMDBClient creates a new TMDB API client
func NewTMDBClient(apiKey string) *TMDBClient {
	return &TMDBClient{
		apikey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
