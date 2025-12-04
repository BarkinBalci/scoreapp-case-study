# scoreapp

A user score calculation service built with Go, featuring clean architecture and automated CI/CD deployment.

[![CI](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/ci.yml/badge.svg)](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/ci.yml)
[![CD](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/cd.yml/badge.svg)](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/cd.yml)
![Go](https://img.shields.io/badge/Go-1.25.4-00ADD8?logo=go)

## Live Demo
 **Production:** [https://scoreapp.barkin.dev](https://scoreapp.barkin.dev)

Try the API:
```bash
# Health check
curl https://scoreapp.barkin.dev/health
```
```bash
# Calculate score
curl -X POST https://scoreapp.barkin.dev/scores/calculate?user_id=user_active
```

## Prerequisites

- Go 1.25+
- Docker (optional)

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