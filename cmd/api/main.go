package main

import (
	"log"
	"net/http"

	"scoreapp/domain"
	"scoreapp/infrastructure/repository"
	httpiface "scoreapp/interfaces/http"
	"scoreapp/usecase"
)

// DummyActionService is a basic example implementation of ActionService.
//
// TODO (candidate):
//   - You can keep this simple and in-memory,
//     or replace it with a more realistic implementation.
type DummyActionService struct{}

func (d *DummyActionService) GetActions(userID string) ([]domain.UserAction, error) {
	// TODO (candidate): return meaningful test data based on userID.
	return nil, nil
}

func main() {
	repo := repository.NewMemoryRepository()
	actionService := &DummyActionService{}

	calculator := usecase.NewScoreCalculator(actionService, repo)
	handler := httpiface.NewScoreHandler(calculator)

	http.HandleFunc("/scores/calculate", handler.CalculateScore)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
