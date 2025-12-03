package repository

import (
	"sync"

	"scoreapp/domain"
)

// MemoryRepository is a simple in-memory example implementation of ScoreRepository.
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
func (r *MemoryRepository) Save(score domain.UserScore) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[score.UserID] = score
	return nil
}

// Get retrieves a score for a given user.
func (r *MemoryRepository) Get(userID string) (domain.UserScore, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	score, exists := r.store[userID]
	return score, exists
}
