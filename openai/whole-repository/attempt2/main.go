package main

import (
	"fmt"
)

func main() {
	board := NewBoard()
	human := "X"
	computer := "O"
	currentPlayer := human

	for {
		fmt.Println(board)
		if winner := board.CheckWinner(); winner != "" {
			fmt.Printf("Player %s wins!\n", winner)
			break
		}
		if board.IsFull() {
			fmt.Println("It's a draw!")
			break
		}

		if currentPlayer == human {
			var row, col int
			fmt.Printf("Player %s, enter row and column (e.g., 1 2): ", human)
			_, err := fmt.Scanf("%d %d", &row, &col)
			if err != nil || !board.MakeMove(human, row, col) {
				fmt.Println("Invalid move, try again.")
				continue
			}
			currentPlayer = computer
		} else {
			row, col := GetBestMove(board, computer)
			board.MakeMove(computer, row, col)
			fmt.Printf("Computer plays at row %d, col %d\n", row, col)
			currentPlayer = human
		}
	}
}
