package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Board represents the Tic Tac Toe board
type Board [9]string

// EmptyBoard initializes an empty board
func EmptyBoard() Board {
	return Board{"e", "e", "e", "e", "e", "e", "e", "e", "e"}
}

// PlayerMove attempts to place a player's mark at the given position
func PlayerMove(board Board, position int, player string) Board {
	if position >= 0 && position < 9 && board[position] == "e" {
		board[position] = player
	}
	return board
}

// FullBoard checks if the board is full
func FullBoard(board Board) bool {
	for _, space := range board {
		if space == "e" {
			return false
		}
	}
	return true
}

// WinningCombinations defines all possible win conditions
var WinningCombinations = [][]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // columns
	{0, 4, 8}, {2, 4, 6}, // diagonals
}

// HasWinner checks if the board has a winner
func HasWinner(board Board) bool {
	for _, combo := range WinningCombinations {
		marks := []string{board[combo[0]], board[combo[1]], board[combo[2]]}
		if marks[0] != "e" && marks[0] == marks[1] && marks[1] == marks[2] {
			return true
		}
	}
	return false
}

// GameOver checks if the game is over (win or draw)
func GameOver(board Board) bool {
	return HasWinner(board) || FullBoard(board)
}

// PrintBoard displays the board in a 3x3 grid
func PrintBoard(board Board) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%s | %s | %s\n", markToString(board[i]), markToString(board[i+1]), markToString(board[i+2]))
		if i < 6 {
			fmt.Println("---------")
		}
	}
}

// markToString converts a mark to a display string
func markToString(mark string) string {
	if mark == "e" {
		return " "
	}
	return mark
}

// getPlayerInput gets a valid move from the player
func getPlayerInput(player string, board Board) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Player %s, enter position (0-8): ", player)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		pos, err := strconv.Atoi(input)
		if err != nil || pos < 0 || pos > 8 {
			fmt.Println("Invalid input. Please enter a number between 0 and 8.")
			continue
		}
		if board[pos] != "e" {
			fmt.Println("Position already taken. Try again.")
			continue
		}
		return pos
	}
}

func main() {
	board := EmptyBoard()
	currentPlayer := "x"

	for !GameOver(board) {
		fmt.Println("\nCurrent board:")
		PrintBoard(board)
		pos := getPlayerInput(currentPlayer, board)
		board = PlayerMove(board, pos, currentPlayer)
		if HasWinner(board) {
			fmt.Println("\nFinal board:")
			PrintBoard(board)
			fmt.Printf("Player %s wins!\n", currentPlayer)
			return
		}
		if FullBoard(board) {
			fmt.Println("\nFinal board:")
			PrintBoard(board)
			fmt.Println("It's a draw!")
			return
		}
		currentPlayer = map[string]string{"x": "o", "o": "x"}[currentPlayer]
	}
}
