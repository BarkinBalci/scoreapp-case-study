# scoreapp

A user score calculation service built with Go, featuring clean architecture and automated CI/CD deployment.

[![CI](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/ci.yml/badge.svg)](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/ci.yml)
[![CD](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/cd.yml/badge.svg)](https://github.com/BarkinBalci/scoreapp-case-study/actions/workflows/cd.yml)
![Go](https://img.shields.io/badge/Go-1.25.4-00ADD8?logo=go)

## Prerequisites

- Go 1.25+
- Docker (optional)

## Quick Start

```bash
# Install dependencies
go mod download
```
```bash
# Start server
make run
```
```bash
# Health check
curl http://localhost:8080/health
```
```bash
# Calculate score
curl -X POST http://localhost:8080/scores/calculate?user_id=user_active
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
