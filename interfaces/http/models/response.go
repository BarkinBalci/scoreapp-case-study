package models

// ScoreResponse represents the response for score calculation endpoints.
type ScoreResponse struct {
	UserID string `json:"user_id"`
	Score  int    `json:"score"`
}

// ErrorResponse represents the response for error cases.
type ErrorResponse struct {
	Error string `json:"error"`
}

// HealthResponse represents the response for health check endpoints.
type HealthResponse struct {
	Status string `json:"status"`
}
