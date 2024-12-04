package display

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"tmdb-cli/pkg/models"
)

type OutputFormat string

const (
	FormatTable OutputFormat = "table"
	FormatJSON  OutputFormat = "json"
)

type MoveOutput struct {
	Format OutputFormat
}

func (mo *MoveOutput) Display(movies []models.Movie) error {
	switch mo.Format {
	case FormatJSON:
		return mo.displayJSON(movies)
	case FormatTable:
		return mo.displayTable(movies)
	default:
		return mo.displayTable(movies)
	}
}

func (mo *MoveOutput) displayJSON(movies []models.Movie) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(movies)
}

func (mo *MoveOutput) displayTable(movies []models.Movie) error {
	fmt.Printf("%-40s | %-15s | %-10s | %-10s\n,", "Title", "Release Date", "Rating", "Popularity")
	fmt.Println(strings.Repeat("-", 80))

	for _, movie := range movies {
		fmt.Printf("%-40s | %-15s | %-10.1f | %-10.2f\n", truncate(movie.Title, 40), movie.ReleaseDate, movie.VoteAverage, movie.Popularity)
	}
	return nil
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}

	return s[:maxLen-3] + "..."
}
