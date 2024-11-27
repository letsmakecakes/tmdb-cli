# TMDB CLI API Integration Guide

## Overview
This CLI tool interacts with The Movie Database (TMDB) API to fetch movie information across different categories.

## API Authentication

### API Key
- Obtain an API key from [TMDB Developer Portal](https://www.themoviedb.org/documentation/api)
- Store the API key in the `.env` file as `TMDB_API_KEY`

### Authentication Method
- Uses API key as a query parameter
- No OAuth or additional authentication required

## Supported Endpoints

### 1. Now Playing Movies
- **Endpoint**: `/movie/now_playing`
- **HTTP Method**: GET
- **Description**: Retrieves movies currently playing in theaters

#### Request Parameters
- `api_key`: Your TMDB API key (required)
- `page`: Page number of results (optional, default: 1)
- `language`: Language of results (optional, default: en-US)

#### Example CLI Usage
```bash
tmdb-cli --type playing
```

### 2. Popular Movies
- **Endpoint**: `/movie/popular`
- **HTTP Method**: GET
- **Description**: Retrieves currently popular movies

#### Request Parameters
- `api_key`: Your TMDB API key (required)
- `page`: Page number of results (optional, default: 1)
- `language`: Language of results (optional, default: en-US)

#### Example CLI Usage
```bash
tmdb-cli --type popular
```

### 3. Top Rated Movies
- **Endpoint**: `/movie/top_rated`
- **HTTP Method**: GET
- **Description**: Retrieves movies with the highest user ratings

#### Request Parameters
- `api_key`: Your TMDB API key (required)
- `page`: Page number of results (optional, default: 1)
- `language`: Language of results (optional, default: en-US)

#### Example CLI Usage
```bash
tmdb-cli --type top
```

### 4. Upcoming Movies
- **Endpoint**: `/movie/upcoming`
- **HTTP Method**: GET
- **Description**: Retrieves movies that are yet to be released

#### Request Parameters
- `api_key`: Your TMDB API key (required)
- `page`: Page number of results (optional, default: 1)
- `language`: Language of results (optional, default: en-US)

#### Example CLI Usage
```bash
tmdb-cli --type upcoming
```

## Response Structure

### Movie Object
```json
{
  "id": integer,
  "title": string,
  "overview": string,
  "release_date": string,
  "vote_average": float,
  "popularity": float,
  "poster_path": string
}
```

### Full Response Example
```json
{
  "page": 1,
  "results": [
    {
      "id": 123456,
      "title": "Example Movie",
      "overview": "A brief description of the movie",
      "release_date": "2024-01-15",
      "vote_average": 7.5,
      "popularity": 450.234,
      "poster_path": "/example_poster.jpg"
    }
  ],
  "total_results": 10000,
  "total_pages": 500
}
```

## Rate Limits and Best Practices

### Rate Limiting
- TMDB typically allows 40 requests per 10 seconds
- Implement exponential backoff for retries
- Cache responses when possible

### Error Handling
Possible HTTP Status Codes:
- `200`: Successful request
- `401`: Invalid API key
- `404`: Resource not found
- `429`: Too many requests

## Potential Improvements
- Support for more detailed movie filtering
- Add support for movie search
- Implement pagination in CLI
- Cache API responses

## Troubleshooting
- Ensure API key is correctly set in `.env`
- Check network connectivity
- Verify TMDB service status at [TMDB Status Page](https://www.themoviedb.org/status)

## API Version
- Current implementation uses TMDB API v3
- Check [TMDB API Documentation](https://developers.themoviedb.org/3/getting-started/introduction) for latest updates

## Security Considerations
- Never commit your API key to version control
- Use environment variables or secure secret management
- Implement rate limiting to prevent abuse