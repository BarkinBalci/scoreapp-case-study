package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"scoreapp/interfaces/http/models"
	"scoreapp/usecase"
)

// ScoreCalculator defines the interface for score calculation.
type ScoreCalculator interface {
	Calculate(userID string) (int, error)
}

// ScoreHandler exposes HTTP endpoints for score calculation.
type ScoreHandler struct {
	calculator ScoreCalculator
}

// NewScoreHandler creates a new ScoreHandler.
func NewScoreHandler(c ScoreCalculator) *ScoreHandler {
	return &ScoreHandler{
		calculator: c,
	}
}

func (h *ScoreHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(models.ErrorResponse{Error: "method not allowed"})
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.ErrorResponse{Error: "user_id is required"})
		return
	}

	score, err := h.calculator.Calculate(userID)
	if err != nil {
		// Check if the error is user not found
		if errors.Is(err, usecase.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(models.ErrorResponse{Error: "user not found"})
			return
		}
		// Other errors are internal server errors
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.ScoreResponse{UserID: userID, Score: score})
}
