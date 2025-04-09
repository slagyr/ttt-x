package game

import (
	"fmt"

	"ttt/board"
	"ttt/player"
)

// PlayGame runs the Tic-Tac-Toe game.
func PlayGame(b *board.Board) string {
	currentPlayer := "X"
	for {
		if currentPlayer == "X" {
			move := player.GetHumanMove(b)
			b.Mark(move, "X")
		} else {
			move := player.GetComputerMove(b)
			b.Mark(move, "O")
			fmt.Printf("Computer chose position %d\n", move)
		}
		winner := b.GetWinner()
		if winner != "" {
			return winner // Return only the winner string
		}
		if b.IsFull() {
			return "" // Return empty string for a draw
		}
		currentPlayer = switchPlayer(currentPlayer)
	}
}

// switchPlayer toggles between "X" and "O".
func switchPlayer(current string) string {
	if current == "X" {
		return "O"
	}
	return "X"
}
