package board

import (
	"fmt"
	"strconv"
)

// Cell represents a cell on the board
type Cell string

// Board is a 3x3 tic-tac-toe board
type Board [9]Cell

// NewBoard creates and returns a new tic-tac-toe board
func NewBoard() Board {
	var b Board
	for i := 0; i < 9; i++ {
		b[i] = Cell(strconv.Itoa(i + 1))
	}
	return b
}

// GetPaths returns all winning paths (rows, columns, diagonals)
func GetPaths(b Board) [][]Cell {
	paths := [][]Cell{
		// Rows
		{b[0], b[1], b[2]},
		{b[3], b[4], b[5]},
		{b[6], b[7], b[8]},
		// Columns
		{b[0], b[3], b[6]},
		{b[1], b[4], b[7]},
		{b[2], b[5], b[8]},
		// Diagonals
		{b[0], b[4], b[8]},
		{b[2], b[4], b[6]},
	}
	return paths
}

// Score evaluates the board and returns the game status
func Score(b Board) string {
	paths := GetPaths(b)

	// Check if X wins
	for _, path := range paths {
		if path[0] == "X" && path[1] == "X" && path[2] == "X" {
			return "X wins!"
		}
	}

	// Check if O wins
	for _, path := range paths {
		if path[0] == "O" && path[1] == "O" && path[2] == "O" {
			return "O wins!"
		}
	}

	// Check for tie
	tie := true
	for _, cell := range b {
		if cell != "X" && cell != "O" {
			tie = false
			break
		}
	}
	if tie {
		return "It's a tie!"
	}

	return "in-progress"
}

// EvaluateBoard assigns numeric scores to board states for the minimax algorithm
func EvaluateBoard(b Board) interface{} {
	paths := GetPaths(b)

	// Check if X wins
	for _, path := range paths {
		if path[0] == "X" && path[1] == "X" && path[2] == "X" {
			return -10
		}
	}

	// Check if O wins
	for _, path := range paths {
		if path[0] == "O" && path[1] == "O" && path[2] == "O" {
			return 10
		}
	}

	// Check for tie
	tie := true
	for _, cell := range b {
		if cell != "X" && cell != "O" {
			tie = false
			break
		}
	}
	if tie {
		return 0
	}

	return "in-progress"
}

// IsNumber checks if a cell contains a number
func IsNumber(cell Cell) bool {
	_, err := strconv.Atoi(string(cell))
	return err == nil
}

// PrintBoard displays the current board state
func PrintBoard(b Board) {
	fmt.Printf("%s %s %s\n", formatCell(b[0]), formatCell(b[1]), formatCell(b[2]))
	fmt.Printf("%s %s %s\n", formatCell(b[3]), formatCell(b[4]), formatCell(b[5]))
	fmt.Printf("%s %s %s\n", formatCell(b[6]), formatCell(b[7]), formatCell(b[8]))
}

// formatCell formats a cell for display
func formatCell(cell Cell) string {
	if cell == "X" || cell == "O" {
		return string(cell)
	}
	// Purple color for numbers
	return fmt.Sprintf("\033[35m%s\033[0m", cell)
}