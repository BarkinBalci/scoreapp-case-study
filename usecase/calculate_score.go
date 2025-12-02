package usecase

import "scoreapp/domain"

// ActionService abstracts an external system that returns user actions.
//
// TODO (candidate):
//   - Implement a concrete ActionService (can be dummy / in-memory or an HTTP client).
type ActionService interface {
	GetActions(userID string) ([]domain.UserAction, error)
}

// ScoreRepository abstracts where we persist the calculated score.
//
// TODO (candidate):
//   - Implement one or more concrete repositories (e.g., in-memory, DB-backed).
type ScoreRepository interface {
	Save(score domain.UserScore) error
}

// ScoreCalculator contains the business logic to calculate and persist scores.
//
// TODO (candidate):
//   - Implement scoring rules:
//     login               -> +1
//     challenge_completed -> +10 * Amount
//     quiz_answer         -> +2 * Amount
//   - Consider making the rules easy to extend.
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

// Calculate loads user actions, computes a score, and persists the result.
//
// TODO (candidate):
//   - Fetch actions via ActionService.
//   - Apply scoring rules.
//   - Call ScoreRepository.Save with the calculated UserScore.
//   - Return appropriate errors.
func (c *ScoreCalculator) Calculate(userID string) error {
	// TODO: implement
	return nil
}
