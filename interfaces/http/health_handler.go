package http

import (
	"encoding/json"
	"net/http"

	"scoreapp/interfaces/http/models"
)

// HealthChecker defines the interface for health checking.
type HealthChecker interface {
	Check() (string, error)
}

// HealthHandler exposes HTTP endpoints for health checks.
type HealthHandler struct {
	checker HealthChecker
}

// NewHealthHandler creates a new HealthHandler.
func NewHealthHandler(c HealthChecker) *HealthHandler {
	return &HealthHandler{
		checker: c,
	}
}

// Handle handles GET /health.
//
// swagger:route GET /health health getHealth
//
// Get application health status
//
//	Responses:
//	  200: healthResponse
//	  405: errorResponse
//	  500: errorResponse
func (h *HealthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(models.ErrorResponse{Error: "method not allowed"})
		return
	}

	status, err := h.checker.Check()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.HealthResponse{Status: status})
}
