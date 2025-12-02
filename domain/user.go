package domain

// UserAction represents a single action performed by a user.
// Candidate can extend this if needed.
type UserAction struct {
	Type   string
	Amount int
}

// UserScore represents the calculated score for a given user.
type UserScore struct {
	UserID string
	Score  int
}
