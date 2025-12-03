package usecase

// HealthChecker performs health checks for the application.
type HealthChecker struct{}

// NewHealthChecker creates a new HealthChecker.
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{}
}

// Check performs a simple health check.
func (h *HealthChecker) Check() (string, error) {
	return "ok", nil
}
