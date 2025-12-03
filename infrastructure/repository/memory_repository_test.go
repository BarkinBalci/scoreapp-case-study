package repository

import (
	"testing"

	"scoreapp/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryRepository(t *testing.T) {
	repo := NewMemoryRepository()

	assert.NotNil(t, repo)
	assert.NotNil(t, repo.store)
	assert.Equal(t, 0, len(repo.store))
}

func TestMemoryRepository_Save_NewScore(t *testing.T) {
	repo := NewMemoryRepository()

	score := domain.UserScore{
		UserID: "user",
		Score:  100,
	}

	err := repo.Save(score)

	assert.NoError(t, err)

	savedScore, exists := repo.Get("user")
	assert.True(t, exists)
	assert.Equal(t, "user", savedScore.UserID)
	assert.Equal(t, 100, savedScore.Score)
}

func TestMemoryRepository_Save_UpdateScore(t *testing.T) {
	repo := NewMemoryRepository()

	initialScore := domain.UserScore{
		UserID: "user",
		Score:  100,
	}
	err := repo.Save(initialScore)
	assert.NoError(t, err)

	updatedScore := domain.UserScore{
		UserID: "user",
		Score:  250,
	}
	err = repo.Save(updatedScore)
	assert.NoError(t, err)

	savedScore, exists := repo.Get("user")
	assert.True(t, exists)
	assert.Equal(t, "user", savedScore.UserID)
	assert.Equal(t, 250, savedScore.Score)
}

func TestMemoryRepository_Get_ExistingUser(t *testing.T) {
	repo := NewMemoryRepository()

	expectedScore := domain.UserScore{
		UserID: "user",
		Score:  75,
	}
	err := repo.Save(expectedScore)
	assert.NoError(t, err)

	actualScore, exists := repo.Get("user")

	assert.True(t, exists)
	assert.Equal(t, expectedScore.UserID, actualScore.UserID)
	assert.Equal(t, expectedScore.Score, actualScore.Score)
}

func TestMemoryRepository_Get_NonExistentUser(t *testing.T) {
	repo := NewMemoryRepository()

	score, exists := repo.Get("nonexistent")

	assert.False(t, exists)
	assert.Equal(t, "", score.UserID)
	assert.Equal(t, 0, score.Score)
}

func TestMemoryRepository_SaveMultipleUsers(t *testing.T) {
	repo := NewMemoryRepository()

	users := []domain.UserScore{
		{UserID: "user1", Score: 10},
		{UserID: "user2", Score: 20},
		{UserID: "user3", Score: 30},
	}

	for _, user := range users {
		err := repo.Save(user)
		assert.NoError(t, err)
	}

	for _, expectedUser := range users {
		actualUser, exists := repo.Get(expectedUser.UserID)
		assert.True(t, exists)
		assert.Equal(t, expectedUser.UserID, actualUser.UserID)
		assert.Equal(t, expectedUser.Score, actualUser.Score)
	}
}

func TestMemoryRepository_SaveEmptyUserID(t *testing.T) {
	repo := NewMemoryRepository()

	score := domain.UserScore{
		UserID: "",
		Score:  100,
	}

	err := repo.Save(score)
	assert.NoError(t, err)

	savedScore, exists := repo.Get("")
	assert.True(t, exists)
	assert.Equal(t, "", savedScore.UserID)
	assert.Equal(t, 100, savedScore.Score)
}

func TestMemoryRepository_SaveZeroScore(t *testing.T) {
	repo := NewMemoryRepository()

	score := domain.UserScore{
		UserID: "user",
		Score:  0,
	}

	err := repo.Save(score)
	assert.NoError(t, err)

	savedScore, exists := repo.Get("user")
	assert.True(t, exists)
	assert.Equal(t, "user", savedScore.UserID)
	assert.Equal(t, 0, savedScore.Score)
}

func TestMemoryRepository_SaveNegativeScore(t *testing.T) {
	repo := NewMemoryRepository()

	score := domain.UserScore{
		UserID: "user",
		Score:  -50,
	}

	err := repo.Save(score)
	assert.NoError(t, err)

	savedScore, exists := repo.Get("user")
	assert.True(t, exists)
	assert.Equal(t, "user", savedScore.UserID)
	assert.Equal(t, -50, savedScore.Score)
}
