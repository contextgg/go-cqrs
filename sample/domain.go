package main

// type Aggregates struct {
// }

// func (*Aggregates) Leaderboard() Aggregate {
// 	return &Leaderboard{}
// }

// Leaderboard aggregate
type Leaderboard struct {
}

// LeaderboardCreated event
type LeaderboardEvents struct {
}

func (*LeaderboardEvents) LeaderboardCreated() {
	// return an event
}
