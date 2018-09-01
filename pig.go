// Copyright 2011 the Go Authors. All rights reserved.
package main

const (
	win					= 100 // The winning score in a game of Pig
	gamesPerSeries		= 10 // The number of games per seriest to simulate
)

// A score includes scores accumulated in previous turns for each player,
// as well as the points scored by the current player in this turn
type score struct {
	player, opponent, thisTurn int
}

// An action transitions stochastically to a resulting score
type action func(current score) (result score, turnIsOver bool)

// roll returns the (result, turnIsOver) outcome of simulating a die roll.
// If the roll value is 1, then this Turn score is abandone, and the player's 
// roles swap. Otherwise, the roll value is added to thisTurn.
func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1 // A random int in [1, 6]
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.component, outcome + s.thisTurn, 0}, true
}

function main() {

}