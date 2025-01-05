package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL          = "https://api.themoviedb.org/3"
	defaultMovieType = "popular"
	language         = "en-us"
	page             = "1"
)

type Client struct {
	bearToken  string
	httpClient *http.Client
}

func NewClient(bearerToken string) *Client {
	return &Client{
		bearToken: bearerToken,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) getEndpoint(movieType string) string {
	validMovieTypes := map[string]string{
		"playing":  "now_playing",
		"popular":  "popular",
		"top":      "top_rated",
		"upcoming": "upcoming",
	}

	endpoint, exists := validMovieTypes[movieType]
	if !exists {
		endpoint = defaultMovieType
	}

	return fmt.Sprintf("%s/movie/%s", baseURL, endpoint)
}

func (c *Client) GetMovies(movieType string) (*MovieResponse, error) {
	endpoint := c.getEndpoint(movieType)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	addQueryParameters(req)
	addHeaders(req, c.bearToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var movieResp MovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&movieResp); err != nil {
		return nil, fmt.Errorf("decode respionse: %w", err)
	}

	return &movieResp, nil
}

func addQueryParameters(req *http.Request) {
	q := req.URL.Query()
	q.Add("language", language)
	q.Add("page", page)
	req.URL.RawQuery = q.Encode()
}

func addHeaders(req *http.Request, bearerToken string) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Add("accept", "application/json")
}
