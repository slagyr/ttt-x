package ui

import (
	"fmt"

	"ttt/board"
	"ttt/game"
)

func StartGame() {
	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("You are X, Computer is O")
	b := board.NewBoard()
	winner := game.PlayGame(b) // Pass the board and expect one return value
	displayBoard(b)            // Display the final board state
	if winner == "X" {
		fmt.Println("You win!")
	} else if winner == "O" {
		fmt.Println("Computer wins!")
	} else {
		fmt.Println("It's a draw!")
	}
}

func displayBoard(b *board.Board) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%s | %s | %s\n", formatCell(b.Cells[i]), formatCell(b.Cells[i+1]), formatCell(b.Cells[i+2]))
		if i < 6 {
			fmt.Println("---------")
		}
	}
	fmt.Println()
}

func formatCell(cell string) string {
	if cell == "" {
		return " "
	}
	return cell
}
