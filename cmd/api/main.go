package main

import (
	"fmt"
	"log"
	"net/http"

	"scoreapp/config"
	"scoreapp/domain"
	"scoreapp/infrastructure/repository"
	httpiface "scoreapp/interfaces/http"
	"scoreapp/usecase"
)

// DummyActionService provides test data based on user ID patterns.
type DummyActionService struct{}

func (d *DummyActionService) GetActions(userID string) ([]domain.UserAction, error) {
	switch userID {
	case "user_beginner":
		return []domain.UserAction{
			{Type: "login", Amount: 1},
			{Type: "challenge_completed", Amount: 0},
			{Type: "quiz_answer", Amount: 0},
		}, nil

	case "user_active":
		return []domain.UserAction{
			{Type: "login", Amount: 1},
			{Type: "challenge_completed", Amount: 2},
			{Type: "quiz_answer", Amount: 3},
		}, nil

	case "user_power":
		return []domain.UserAction{
			{Type: "login", Amount: 0},
			{Type: "challenge_completed", Amount: 10},
			{Type: "quiz_answer", Amount: 25},
		}, nil

	case "user_empty":
		return []domain.UserAction{}, nil

	case "user_error":
		return nil, fmt.Errorf("simulated service error")

	default:
		return nil, usecase.ErrUserNotFound
	}
}

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the repository based on configuration
	repo := repository.NewMemoryRepository()

	// Initialize services
	actionService := &DummyActionService{}
	calculator := usecase.NewScoreCalculator(actionService, repo)

	// Initialize health checker
	healthChecker := usecase.NewHealthChecker()

	// Initialize handlers
	scoreHandler := httpiface.NewScoreHandler(calculator)
	healthHandler := httpiface.NewHealthHandler(healthChecker)

	// Register routes
	http.HandleFunc("/scores/calculate", scoreHandler.Handle)
	http.HandleFunc("/health", healthHandler.Handle)

	// Start server
	addr := ":" + cfg.Server.Port
	log.Printf("Starting server on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
