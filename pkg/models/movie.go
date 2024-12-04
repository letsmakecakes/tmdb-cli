package models

type Movie struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Overview      string  `json:"overview"`
	ReleaseDate   string  `json:"release_date"`
	Poster        string  `json:"poster_path"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
	Popularity    float64 `json:"popularity"`
	OriginalTitle string  `json:"original_title"`
	Language      string  `json:"original_language"`
	Adult         bool    `json:"adult"`
	Genres        []Genre `json:"genres,omitempty"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MovieResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}
