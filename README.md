# Go Configuration Library

A shared library for configuration management across FileConvert microservices.

## Features

- Environment variable management
- Configuration file loading
- Configuration validation
- Default values
- Type-safe configuration

## Usage

```go
import "github.com/file-convert/go-config"

// Create a new configuration
config := goconfig.NewConfig()

// Load configuration from file
err := config.LoadFromFile("config.yaml")
if err != nil {
    log.Fatal(err)
}

// Get configuration values
dbUri := config.GetString("DB_URI")
port := config.GetInt("APP_PORT")

// Set configuration values
config.Set("KEY", "value")

// Validate configuration
err = config.Validate()
if err != nil {
    log.Fatal(err)
}
```

## Configuration

The library supports the following configuration options:

- File formats (YAML, JSON, ENV)
- Environment variable prefix
- Required fields
- Default values
- Validation rules 