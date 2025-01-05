# TMDB CLI Tool

A command-line interface (CLI) tool built in Go that fetches and displays movie information from The Movie Database (TMDB) API. Get quick access to popular, top-rated, upcoming, and now playing movies right from your terminal.

## 🎯 Features

- Fetch different categories of movies:
    - Now Playing
    - Popular Movies
    - Top Rated Movies
    - Upcoming Movies
- Clean terminal output with formatted movie details
- Error handling for API and network issues
- Configurable via environment variables
- Simple and intuitive command-line interface

## 🚀 Quick Start

### Prerequisites

- Go 1.16 or higher
- TMDB Bearer Token ([Get it here](https://www.themoviedb.org/settings/api))

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/tmdb-app.git
cd tmdb-app
```

2. Set up your TMDB Bearer Token:
```bash
export TMDB_BEARER_TOKEN='your-bearer-token-here'
```

3. Build the application:
```bash
make build
```

### Usage

The CLI tool supports the following commands:

```bash
# Get now playing movies
./bin/tmdb-app --type "playing"

# Get popular movies
./bin/tmdb-app --type "popular"

# Get top rated movies
./bin/tmdb-app --type "top"

# Get upcoming movies
./bin/tmdb-app --type "upcoming"
```

## 📁 Project Structure

```
.
├── cmd/
│   └── tmdb-app/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── client.go
│   │   └── models.go
│   ├── config/
│   │   └── config.go
│   └── cli/
│       └── app.go
├── pkg/
│   └── formatter/
│       └── formatter.go
├── go.mod
├── Makefile
└── README.md
```

### Directory Overview

- `cmd/`: Contains the main application entry points
- `internal/`: Private application code
    - `api/`: TMDB API client and data models
    - `config/`: Configuration management
    - `cli/`: CLI application logic
- `pkg/`: Public packages that can be used by external applications
    - `formatter/`: Output formatting utilities

## 🛠️ Development

### Building

```bash
# Build the application
make build

# Run tests
make test

# Run the application directly
make run
```

### Adding New Features

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 Example Output

```
Movie List:
====================================

Title: Gladiator II (2024-11-05)
Rating: 6.8/10 (1979 votes)
Popularity: 4782.6

Overview:
Years after witnessing the death of the revered hero Maximus...
------------------------------------

Title: Your Fault (2024-12-26)
Rating: 7.1/10 (641 votes)
Popularity: 4279.1

Overview:
The love between Noah and Nick seems unwavering...
------------------------------------

Total movies shown: 20
```

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Description | Required |
|----------|-------------|----------|
| TMDB_BEARER_TOKEN | Your TMDB API Bearer Token | Yes |

## ⚠️ Error Handling

The application handles various error cases:
- Invalid movie types
- API authentication failures
- Network connectivity issues
- Rate limiting
- Invalid responses

## 📚 API Documentation

This project uses the TMDB API v3. For more information about the API, visit:
- [TMDB API Documentation](https://developers.themoviedb.org/3)
- [Authentication](https://developers.themoviedb.org/3/authentication)
- [Movie Endpoints](https://developers.themoviedb.org/3/movies)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [The Movie Database (TMDB)](https://www.themoviedb.org/) for providing the API
- The Go community for the amazing tooling and libraries

## 📧 Contact

Your Name - [@letsmakecakes](https://twitter.com/letsmakecakes)

Project Link: [https://github.com/letsmakecakes/tmdb-cli](https://github.com/letsmakecakes/tmdb-clie)