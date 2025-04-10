// Package tic_tac_toe provides functions for evaluating the Tic-Tac-Toe board state.
package tic_tac_toe

// Board represents the Tic-Tac-Toe board.
type Board [][]string

// CheckWin evaluates whether the given player has won the game.
func CheckWin(board Board, player string) bool {
	// Check rows and columns
	for i := 0; i < len(board); i++ {
		if checkLine(board[i], player) || checkLine(getColumn(board, i), player) {
			return true
		}
	}

	// Check diagonals
	return checkLine(getDiagonal(board, true), player) || checkLine(getDiagonal(board, false), player)
}

// checkLine checks if all elements in a line belong to the specified player.
func checkLine(line []string, player string) bool {
	for _, v := range line {
		if v != player {
			return false
		}
	}
	return true
}

// getColumn returns a specific column from the board.
func getColumn(board Board, index int) []string {
	column := make([]string, len(board))
	for i, row := range board {
		column[i] = row[index]
	}
	return column
}

// getDiagonal returns one of the two diagonals from the board.
func getDiagonal(board Board, main bool) []string {
	diagonal := make([]string, len(board))
	for i := 0; i < len(board); i++ {
		if main {
			diagonal[i] = board[i][i]
		} else {
			diagonal[i] = board[i][len(board)-i-1]
		}
	}
	return diagonal
}
