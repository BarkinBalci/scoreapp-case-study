package repository

import (
	"sync"

	"scoreapp/domain"
)

// MemoryRepository is a simple in-memory example implementation of ScoreRepository.
//
// TODO (candidate):
//   - Ensure this satisfies usecase.ScoreRepository.
//   - Optionally add helper methods for tests (e.g. Get).
type MemoryRepository struct {
	mu    sync.Mutex
	store map[string]domain.UserScore
}

// NewMemoryRepository creates a new MemoryRepository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		store: make(map[string]domain.UserScore),
	}
}

// Save stores or updates the score for a given user.
//
// TODO (candidate):
//   - Implement this so that it writes to the in-memory map.
func (r *MemoryRepository) Save(score domain.UserScore) error {
	// TODO: implement
	return nil
}
