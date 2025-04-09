package main

import (
	"fmt"
)

func main() {
	board := NewBoard()
	player := "X"

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

		var row, col int
		fmt.Printf("Player %s, enter row and column (e.g., 1 2): ", player)
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil || !board.MakeMove(player, row, col) {
			fmt.Println("Invalid move, try again.")
			continue
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}
