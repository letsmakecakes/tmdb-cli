package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"tmdb-cli/pkg/models"
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

// FetchMovies retrieves movies based on the specified type
func (c *TMDBClient) FetchMovies(movieType string) ([]models.Movie, error) {
	endpoint := c.getEndpoint(movieType)
	if endpoint == "" {
		return nil, errors.New("invalid movie type: " + movieType)
	}

	url := fmt.Sprintf("%s%s?api_key=%s", baseURL, endpoint, c.apikey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, errors.New("failed to fetch movies: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned code %d", resp.StatusCode)
	}

	var movieResponse models.MovieResponse

	if err := json.NewDecoder(resp.Body).Decode(&movieResponse); err != nil {
		return nil, errors.New("failed to decode response: " + err.Error())
	}

	return movieResponse.Results, nil
}

// getEndpoint maps movie type to TMDB API endpoint
func (c *TMDBClient) getEndpoint(movieType string) string {
	switch movieType {
	case "playing":
		return "/movie/now_playing"
	case "popular":
		return "/movie/popular"
	case "top":
		return "/movie/top_rated"
	case "upcoming":
		return "/movie/upcoming"
	default:
		return ""
	}
}
