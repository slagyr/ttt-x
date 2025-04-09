package main

import (
	"fmt"
	"tic-tac-toe/board"
	"tic-tac-toe/player"
)

func main() {
	// Initialize the board
	b := board.NewBoard()
	currentPlayer := "X" // Player X starts

	// Game loop
	for {
		// Print the board
		board.PrintBoard(b)

		// Check if game is over
		score := board.Score(b)
		if score != "in-progress" {
			fmt.Println(score)
			break
		}

		// Play turn based on current player
		if currentPlayer == "X" {
			b = player.PlayUserTurn(b, currentPlayer)
		} else {
			b = player.PlayBotTurn(b, currentPlayer)
		}

		// Switch player
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}