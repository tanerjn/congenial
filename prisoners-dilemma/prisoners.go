package main

import (
	"fmt"
)

// Player represents a prisoner making decisions.
type Player struct {
	name     string
	strategy string // "C" for cooperate, "D" for defect
}

// Payoff matrix
var payoffMatrix = map[string]map[string]int{
	"C": {"C": -1, "D": -3},
	"D": {"C": 0, "D": -2},
}

// play simulates a round of the Prisoner's Dilemma between two players.
func play(player1, player2 Player) {
	fmt.Printf("%s chooses to %s, %s chooses to %s\n", player1.name, player1.strategy, player2.name, player2.strategy)

	// Get the payoffs based on their strategies
	player1Payoff := payoffMatrix[player1.strategy][player2.strategy]
	player2Payoff := payoffMatrix[player2.strategy][player1.strategy]

	// Display the outcome
	fmt.Printf("Outcomes: %s: %d years, %s: %d years\n\n", player1.name, player1Payoff, player2.name, player2Payoff)
}

func main() {
	// Define two players
	player1 := Player{name: "Player 1", strategy: "C"} // Player 1 cooperates
	player2 := Player{name: "Player 2", strategy: "D"} // Player 2 defects

	// Play the game
	play(player1, player2)

	// You can change the strategies and play again
	player1.strategy = "D" // Now Player 1 defects
	player2.strategy = "C" // Player 2 cooperates
	play(player1, player2)

	// Another round with both players defecting
	player1.strategy = "D"
	player2.strategy = "D"
	play(player1, player2)
}
