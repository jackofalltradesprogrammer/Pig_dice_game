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

function main() {

}