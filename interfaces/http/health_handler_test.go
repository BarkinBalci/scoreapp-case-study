package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"scoreapp/interfaces/http/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHealthChecker is a mock for HealthChecker.
type MockHealthChecker struct {
	mock.Mock
}

func (m *MockHealthChecker) Check() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestHealthHandle_Success(t *testing.T) {
	mockChecker := new(MockHealthChecker)
	handler := NewHealthHandler(mockChecker)

	mockChecker.On("Check").Return("ok", nil)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.HealthResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response.Status)

	mockChecker.AssertExpectations(t)
}

func TestHealthHandle_MethodNotAllowed(t *testing.T) {
	mockChecker := new(MockHealthChecker)
	handler := NewHealthHandler(mockChecker)

	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "method not allowed", response.Error)

	mockChecker.AssertNotCalled(t, "Check")
}

func TestHealthHandle_Error(t *testing.T) {
	mockChecker := new(MockHealthChecker)
	handler := NewHealthHandler(mockChecker)

	checkError := errors.New("health check failed")
	mockChecker.On("Check").Return("", checkError)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.Handle(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, checkError.Error(), response.Error)

	mockChecker.AssertExpectations(t)
}
