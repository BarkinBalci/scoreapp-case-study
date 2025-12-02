package usecase

import (
	"testing"

	"scoreapp/domain"

	"github.com/stretchr/testify/mock"
)

// MockActionService is a testify mock for ActionService.
type MockActionService struct {
	mock.Mock
}

func (m *MockActionService) GetActions(userID string) ([]domain.UserAction, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.UserAction), args.Error(1)
}

// MockScoreRepository is a testify mock for ScoreRepository.
type MockScoreRepository struct {
	mock.Mock
}

func (m *MockScoreRepository) Save(score domain.UserScore) error {
	args := m.Called(score)
	return args.Error(0)
}

// TODO (candidate):
//   - Replace this placeholder test with real tests that check scoring logic.
//   - Cover happy path and at least one error path.
func TestScoreCalculation_Placeholder(t *testing.T) {
	t.Skip("candidate should implement ScoreCalculator tests using mocks")
}
