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

// PlayerMove places a human player's mark
func PlayerMove(board Board, position int, player string) Board {
	if position >= 0 && position < 9 && board[position] == "e" {
		board[position] = player
	}
	return board
}

// ComputerMove places the computer's mark in the first empty space
func ComputerMove(board Board, player string) Board {
	for i := 0; i < 9; i++ {
		if board[i] == "e" {
			board[i] = player
			return board
		}
	}
	return board // Shouldn't reach here if board isn't full
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

// WinningCombinations defines win conditions
var WinningCombinations = [][]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // columns
	{0, 4, 8}, {2, 4, 6}, // diagonals
}

// HasWinner checks for a winner
func HasWinner(board Board) bool {
	for _, combo := range WinningCombinations {
		marks := []string{board[combo[0]], board[combo[1]], board[combo[2]]}
		if marks[0] != "e" && marks[0] == marks[1] && marks[1] == marks[2] {
			return true
		}
	}
	return false
}

// GameOver checks if the game is over
func GameOver(board Board) bool {
	return HasWinner(board) || FullBoard(board)
}

// PrintBoard displays the board with position numbers for empty spaces
func PrintBoard(board Board) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%s | %s | %s\n", markToString(board[i], i), markToString(board[i+1], i+1), markToString(board[i+2], i+2))
		if i < 6 {
			fmt.Println("---------")
		}
	}
}

// markToString shows position number if empty, otherwise the mark
func markToString(mark string, pos int) string {
	if mark == "e" {
		return fmt.Sprintf("%d", pos)
	}
	return mark
}

// getPlayerInput gets a valid move from the human player
func getPlayerInput(board Board) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Your move (0-8): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		pos, err := strconv.Atoi(input)
		if err != nil || pos < 0 || pos > 8 {
			fmt.Println("Invalid input. Enter a number between 0 and 8.")
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
	human := "x"
	computer := "o"

	fmt.Println("Welcome to Tic Tac Toe! You are X, computer is O.")
	for !GameOver(board) {
		// Human turn
		fmt.Println("\nCurrent board:")
		PrintBoard(board)
		pos := getPlayerInput(board)
		board = PlayerMove(board, pos, human)
		if HasWinner(board) {
			fmt.Println("\nFinal board:")
			PrintBoard(board)
			fmt.Println("You win!")
			return
		}
		if FullBoard(board) {
			fmt.Println("\nFinal board:")
			PrintBoard(board)
			fmt.Println("It's a draw!")
			return
		}

		// Computer turn
		fmt.Println("\nComputer's turn:")
		board = ComputerMove(board, computer)
		fmt.Println("Current board:")
		PrintBoard(board)
		if HasWinner(board) {
			fmt.Println("Computer wins!")
			return
		}
		if FullBoard(board) {
			fmt.Println("It's a draw!")
			return
		}
	}
}
