// main.go - Entry point for the Tic-Tac-Toe game.
package tic_tac_toe

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	board := tic_tac_toe.NewBoard()
	players := []string{"X", "O"}
	currentPlayer := 0

	for {
		// Print the current board state
		tic_tac_toe.PrintBoard(board)

		// Get player move
		fmt.Printf("Player %s, enter your move (row and column): ", players[currentPlayer])

		x, y := getPlayerMove()

		if tic_tac_toe.MakeMove(board, x, y, players[currentPlayer]) {
			if tic_tac_toe.CheckWin(board, players[currentPlayer]) {
				fmt.Printf("Player %s wins!\n", players[currentPlayer])
				break
			}
			currentPlayer = (currentPlayer + 1) % 2
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

// getPlayerMove reads and returns validated player input for moves
func getPlayerMove() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		var x, y int
		_, err := fmt.Sscanf(input, "%d %d", &x, &y)
		if err != nil {
			fmt.Println("Invalid input. Please enter two numbers separated by a space.")
			continue
		}
		return x, y
	}
}
