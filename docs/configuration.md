# TMDB CLI Configuration Guide

## Overview

The TMDB CLI tool supports multiple configuration methods to provide flexibility and ease of use. This document outlines the various ways you can configure the application.

## Configuration Sources

The application supports multiple configuration sources in the following priority order:

1. Command-line Flags
2. Environment Variables
3. Configuration File
4. Default Values

### 1. API Key Configuration

#### Environment Variable
The primary method to configure your TMDB API key is through an environment variable:

```bash
export TMDB_API_KEY='your_api_key_here'
```

#### .env File
You can also use a `.env` file in the project root:

```
TMDB_API_KEY=your_api_key_here
```

#### Command-line Flag (Not Recommended for Security)
While possible, passing the API key directly via CLI is not recommended:

```bash
tmdb-cli --api-key your_api_key_here
```

## Configuration Options

### API Configuration

| Option          | Environment Variable   | CLI Flag         | Default                        | Description          |
|-----------------|------------------------|------------------|--------------------------------|----------------------|
| API Key         | `TMDB_API_KEY`         | `--api-key`      | None (Required)                | Your TMDB API key    |
| API Endpoint    | `TMDB_API_ENDPOINT`    | `--api-endpoint` | `https://api.themoviedb.org/3` | Base API endpoint    |
| Request Timeout | `TMDB_REQUEST_TIMEOUT` | `--timeout`      | 30 seconds                     | HTTP request timeout |

### Display Configuration

| Option        | Environment Variable | CLI Flag        | Default | Description                          |
|---------------|----------------------|-----------------|---------|--------------------------------------|
| Max Results   | `TMDB_MAX_RESULTS`   | `--max-results` | 20      | Maximum number of movies to display  |
| Color Output  | `TMDB_COLOR_OUTPUT`  | `--color`       | true    | Enable/disable colored output        |
| Detailed View | `TMDB_DETAILED_VIEW` | `--detailed`    | false   | Show more detailed movie information |

## Example Configurations

### Using Environment Variables

```bash
# Set API key
export TMDB_API_KEY='your_key_here'

# Customize display
export TMDB_MAX_RESULTS=50
export TMDB_COLOR_OUTPUT=false

# Run the CLI
tmdb-cli --type popular
```

### Using .env File

Create a `.env` file in your project root:

```
# TMDB API Configuration
TMDB_API_KEY=your_api_key_here
TMDB_API_ENDPOINT=https://api.themoviedb.org/3
TMDB_REQUEST_TIMEOUT=45

# Display Configuration
TMDB_MAX_RESULTS=30
TMDB_COLOR_OUTPUT=true
```

### Command-line Flags

```bash
# Fetch popular movies with custom configuration
tmdb-cli --type popular --max-results 40 --color false
```

## Security Considerations

1. **Never commit API keys to version control**
2. Use `.env` file and add it to `.gitignore`
3. Use environment-specific configurations
4. Rotate API keys periodically

## Troubleshooting

### Common Configuration Errors

- **Missing API Key**:
    - Ensure `TMDB_API_KEY` is set
    - Check API key validity on TMDB website

- **Connection Issues**:
    - Verify internet connection
    - Check API endpoint
    - Increase request timeout if needed

## Best Practices

1. Store sensitive information in environment variables
2. Use a `.env.example` file as a template for configuration
3. Never hard-code API keys in source code
4. Implement proper error handling for configuration loading

## Extending Configuration

The application is designed to be flexible. Developers can extend the configuration management by:
- Adding new configuration sources
- Implementing custom configuration validators
- Creating more complex configuration hierarchies

## Versions

- **Current Version**: 1.0.0
- **Last Updated**: [Current Date]

---

**Note**: Always refer to the most recent documentation for the most up-to-date configuration instructions.