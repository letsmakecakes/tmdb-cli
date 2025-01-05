package formatter

import (
	"fmt"
	"strings"
	"tmdb-cli/internal/api"
)

// PrintMovies prints a formatted list of movies
func PrintMovies(movies []api.Movie) {
	lineSeparator := strings.Repeat("=", 100)
	subSeparator := strings.Repeat("-", 100)

	fmt.Println("\nMovie List:")
	fmt.Println(lineSeparator)

	for _, movie := range movies {
		printMovieDetails(movie, subSeparator)
	}

	fmt.Printf("\nTotal movies shown: %d\n", len(movies))
}

// printMovieDetails prints the details of a single movie
func printMovieDetails(movie api.Movie, separator string) {
	fmt.Printf("\nTitle: %s (%s)\n", movie.Title, movie.ReleaseDate)
	fmt.Printf("Rating: %.1f/10 (%d votes)\n", movie.VoteAverage, movie.VoteCount)
	fmt.Printf("Popularity: %.1f\n", movie.Popularity)
	fmt.Println("\nOverview:")
	fmt.Println(strings.TrimSpace(movie.Overview))
	fmt.Println(separator)
}
