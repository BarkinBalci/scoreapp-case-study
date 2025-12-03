package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"scoreapp/interfaces/http/models"
	"scoreapp/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockScoreCalculator is a mock for ScoreCalculator.
type MockScoreCalculator struct {
	mock.Mock
}

func (m *MockScoreCalculator) Calculate(userID string) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func TestHandle_MethodNotAllowed(t *testing.T) {
	mockCalculator := new(MockScoreCalculator)
	handler := NewScoreHandler(mockCalculator)

	req := httptest.NewRequest(http.MethodGet, "/scores/calculate?user_id=user123", nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "method not allowed", response.Error)

	mockCalculator.AssertNotCalled(t, "Calculate", mock.Anything)
}

func TestHandle_MissingUserID(t *testing.T) {
	mockCalculator := new(MockScoreCalculator)
	handler := NewScoreHandler(mockCalculator)

	req := httptest.NewRequest(http.MethodPost, "/scores/calculate", nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "user_id is required", response.Error)

	mockCalculator.AssertNotCalled(t, "Calculate", mock.Anything)
}

func TestHandle_SuccessfulCalculation(t *testing.T) {
	mockCalculator := new(MockScoreCalculator)
	handler := NewScoreHandler(mockCalculator)

	userID := "user"
	expectedScore := 42

	mockCalculator.On("Calculate", userID).Return(expectedScore, nil)

	req := httptest.NewRequest(http.MethodPost, "/scores/calculate?user_id="+userID, nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ScoreResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, userID, response.UserID)
	assert.Equal(t, expectedScore, response.Score)

	mockCalculator.AssertExpectations(t)
}

func TestHandle_UserNotFound(t *testing.T) {
	mockCalculator := new(MockScoreCalculator)
	handler := NewScoreHandler(mockCalculator)

	userID := "user"

	mockCalculator.On("Calculate", userID).Return(0, usecase.ErrUserNotFound)

	req := httptest.NewRequest(http.MethodPost, "/scores/calculate?user_id="+userID, nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "user not found", response.Error)

	mockCalculator.AssertExpectations(t)
}

func TestHandle_InternalServerError(t *testing.T) {
	mockCalculator := new(MockScoreCalculator)
	handler := NewScoreHandler(mockCalculator)

	userID := "user"
	internalError := errors.New("database connection failed")

	mockCalculator.On("Calculate", userID).Return(0, internalError)

	req := httptest.NewRequest(http.MethodPost, "/scores/calculate?user_id="+userID, nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, internalError.Error(), response.Error)

	mockCalculator.AssertExpectations(t)
}

func TestHandle_EmptyUserID(t *testing.T) {
	mockCalculator := new(MockScoreCalculator)
	handler := NewScoreHandler(mockCalculator)

	req := httptest.NewRequest(http.MethodPost, "/scores/calculate?user_id=", nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "user_id is required", response.Error)

	mockCalculator.AssertNotCalled(t, "Calculate", mock.Anything)
}

func TestHandle_ScoreZero(t *testing.T) {
	mockCalculator := new(MockScoreCalculator)
	handler := NewScoreHandler(mockCalculator)

	userID := "user"
	expectedScore := 0

	mockCalculator.On("Calculate", userID).Return(expectedScore, nil)

	req := httptest.NewRequest(http.MethodPost, "/scores/calculate?user_id="+userID, nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ScoreResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, userID, response.UserID)
	assert.Equal(t, expectedScore, response.Score)

	mockCalculator.AssertExpectations(t)
}
