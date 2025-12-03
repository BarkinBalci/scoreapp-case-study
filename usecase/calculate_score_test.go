package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"scoreapp/domain"
)

// MockActionService is a mock for ActionService.
type MockActionService struct {
	mock.Mock
}

func (m *MockActionService) GetActions(userID string) ([]domain.UserAction, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.UserAction), args.Error(1)
}

// MockScoreRepository is a mock for ScoreRepository.
type MockScoreRepository struct {
	mock.Mock
}

func (m *MockScoreRepository) Save(score domain.UserScore) error {
	args := m.Called(score)
	return args.Error(0)
}

func TestScoreCalculation_HappyPath(t *testing.T) {
	mockActionService := new(MockActionService)
	mockRepo := new(MockScoreRepository)

	userID := "user"
	actions := []domain.UserAction{
		{Type: "login", Amount: 1},
		{Type: "challenge_completed", Amount: 3},
		{Type: "quiz_answer", Amount: 5},
	}
	expectedScore := 41

	mockActionService.On("GetActions", userID).Return(actions, nil)
	mockRepo.On("Save", domain.UserScore{
		UserID: userID,
		Score:  expectedScore,
	}).Return(nil)

	calculator := NewScoreCalculator(mockActionService, mockRepo)

	score, err := calculator.Calculate(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedScore, score)
	mockActionService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestScoreCalculation_EmptyActions(t *testing.T) {
	mockActionService := new(MockActionService)
	mockRepo := new(MockScoreRepository)

	userID := "user"
	var actions []domain.UserAction
	expectedScore := 0

	mockActionService.On("GetActions", userID).Return(actions, nil)
	mockRepo.On("Save", domain.UserScore{
		UserID: userID,
		Score:  expectedScore,
	}).Return(nil)

	calculator := NewScoreCalculator(mockActionService, mockRepo)

	score, err := calculator.Calculate(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedScore, score)
	mockActionService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestScoreCalculation_ActionsWithZeroAmounts(t *testing.T) {
	mockActionService := new(MockActionService)
	mockRepo := new(MockScoreRepository)

	userID := "user"
	actions := []domain.UserAction{
		{Type: "login", Amount: 0},
		{Type: "challenge_completed", Amount: 0},
		{Type: "quiz_answer", Amount: 0},
	}
	expectedScore := 0

	mockActionService.On("GetActions", userID).Return(actions, nil)
	mockRepo.On("Save", domain.UserScore{
		UserID: userID,
		Score:  expectedScore,
	}).Return(nil)

	calculator := NewScoreCalculator(mockActionService, mockRepo)

	score, err := calculator.Calculate(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedScore, score)
	mockActionService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestScoreCalculation_ActionsWithNegativeAmounts(t *testing.T) {
	mockActionService := new(MockActionService)
	mockRepo := new(MockScoreRepository)

	userID := "user"
	actions := []domain.UserAction{
		{Type: "login", Amount: -2},
		{Type: "challenge_completed", Amount: -5},
		{Type: "quiz_answer", Amount: -7},
	}
	expectedScore := 0

	mockActionService.On("GetActions", userID).Return(actions, nil)
	mockRepo.On("Save", domain.UserScore{
		UserID: userID,
		Score:  expectedScore,
	}).Return(nil)

	calculator := NewScoreCalculator(mockActionService, mockRepo)

	score, err := calculator.Calculate(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedScore, score)
	mockActionService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestScoreCalculation_UnknownActionType(t *testing.T) {
	mockActionService := new(MockActionService)
	mockRepo := new(MockScoreRepository)

	userID := "user"
	actions := []domain.UserAction{
		{Type: "login", Amount: 1},
		{Type: "unknown_action", Amount: 100}, // Should be ignored
		{Type: "challenge_completed", Amount: 2},
	}
	expectedScore := 21

	mockActionService.On("GetActions", userID).Return(actions, nil)
	mockRepo.On("Save", domain.UserScore{
		UserID: userID,
		Score:  expectedScore,
	}).Return(nil)

	calculator := NewScoreCalculator(mockActionService, mockRepo)

	score, err := calculator.Calculate(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedScore, score)
	mockActionService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestScoreCalculation_ActionServiceError(t *testing.T) {
	mockActionService := new(MockActionService)
	mockRepo := new(MockScoreRepository)

	userID := "user"
	expectedError := errors.New("service unavailable")

	mockActionService.On("GetActions", userID).Return([]domain.UserAction(nil), expectedError)

	calculator := NewScoreCalculator(mockActionService, mockRepo)

	score, err := calculator.Calculate(userID)

	assert.Error(t, err)
	assert.Equal(t, 0, score)
	assert.Contains(t, err.Error(), "failed to get actions")
	mockActionService.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Save", mock.Anything)
}

func TestScoreCalculation_RepositoryError(t *testing.T) {
	mockActionService := new(MockActionService)
	mockRepo := new(MockScoreRepository)

	userID := "user"
	actions := []domain.UserAction{
		{Type: "login", Amount: 0},
	}
	expectedError := errors.New("database error")

	mockActionService.On("GetActions", userID).Return(actions, nil)
	mockRepo.On("Save", mock.Anything).Return(expectedError)

	calculator := NewScoreCalculator(mockActionService, mockRepo)

	score, err := calculator.Calculate(userID)

	assert.Error(t, err)
	assert.Equal(t, 0, score)
	assert.Contains(t, err.Error(), "failed to save score")
	mockActionService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
