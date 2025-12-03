# scoreapp

A user score calculation service.

![Go](https://img.shields.io/badge/Go-1.25.4-00ADD8?logo=go)

## Prerequisites

- Go 1.25+

## Quick Start

```bash
# Install dependencies
go mod download

# Start server
make run
```

## API Documentation

API documentation is available in [docs/swagger.yaml](docs/swagger.yaml).

## Testing

```bash
# Run all tests
make test

# Generate coverage report
make cover
```

## Development

### Running Tests with Coverage

```bash
make cover
open cover.html
```

### Building the Application

```bash
make build
./bin/scoreapp
```