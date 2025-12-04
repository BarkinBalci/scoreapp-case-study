package usecase

import (
	"errors"
	"fmt"

	"scoreapp/domain"
)

// ErrUserNotFound is returned when a user is not found.
var ErrUserNotFound = errors.New("user not found")

// ActionService abstracts an external system that returns user actions.
type ActionService interface {
	GetActions(userID string) ([]domain.UserAction, error)
}

// ScoreRepository abstracts where we persist the calculated score.
type ScoreRepository interface {
	Save(score domain.UserScore) error
}

// ScoreCalculator contains the business logic to calculate and persist scores.
type ScoreCalculator struct {
	actionService ActionService
	repo          ScoreRepository
}

// NewScoreCalculator constructs a ScoreCalculator with its dependencies.
func NewScoreCalculator(a ActionService, r ScoreRepository) *ScoreCalculator {
	return &ScoreCalculator{
		actionService: a,
		repo:          r,
	}
}

// Calculate loads user actions and calculates a score
func (c *ScoreCalculator) Calculate(userID string) (int, error) {
	// Fetch actions from ActionService
	actions, err := c.actionService.GetActions(userID)
	if err != nil {
		return 0, fmt.Errorf("failed to get actions: %w", err)
	}

	// Calculate score based on rules
	score := 0
	for _, action := range actions {
		if action.Amount <= 0 {
			continue
		}

		switch action.Type {
		case "login":
			score += 1
		case "challenge_completed":
			score += 10 * action.Amount
		case "quiz_answer":
			score += 2 * action.Amount
		default:
			continue
		}
	}

	// Create UserScore domain object
	userScore := domain.UserScore{
		UserID: userID,
		Score:  score,
	}

	// Save via repository
	if err := c.repo.Save(userScore); err != nil {
		return 0, fmt.Errorf("failed to save score: %w", err)
	}

	return score, nil
}
