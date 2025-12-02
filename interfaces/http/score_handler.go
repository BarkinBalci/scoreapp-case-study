package http

import (
	"encoding/json"
	"net/http"

	"scoreapp/usecase"
)

// ScoreHandler exposes HTTP endpoints for score calculation.
type ScoreHandler struct {
	calculator *usecase.ScoreCalculator
}

// NewScoreHandler creates a new ScoreHandler.
func NewScoreHandler(c *usecase.ScoreCalculator) *ScoreHandler {
	return &ScoreHandler{
		calculator: c,
	}
}

// CalculateScore handles POST /scores/calculate?user_id=<id>.
//
// TODO (candidate):
//   - Keep handler thin (no business logic here).
//   - Validate input.
//   - Call calculator.Calculate.
//   - Return appropriate HTTP status codes.
func (h *ScoreHandler) CalculateScore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	if err := h.calculator.Calculate(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
