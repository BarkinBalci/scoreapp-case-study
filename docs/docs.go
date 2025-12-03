// Package docs scoreapp API
//
// # User score calculation service
//
// Schemes: http, https
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package docs

import (
	"scoreapp/interfaces/http/models"
)

// swagger:response scoreResponse
//
//nolint:unused
type scoreResponseWrapper struct {
	// in: body
	Body models.ScoreResponse
}

// swagger:response healthResponse
//
//nolint:unused
type healthResponseWrapper struct {
	// in: body
	Body models.HealthResponse
}

// swagger:response errorResponse
//
//nolint:unused
type errorResponseWrapper struct {
	// in: body
	Body models.ErrorResponse
}
