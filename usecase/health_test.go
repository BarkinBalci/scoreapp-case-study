package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHealthChecker(t *testing.T) {
	checker := NewHealthChecker()
	assert.NotNil(t, checker)
}

func TestHealthChecker_Check(t *testing.T) {
	checker := NewHealthChecker()

	status, err := checker.Check()

	assert.NoError(t, err)
	assert.Equal(t, "ok", status)
}
