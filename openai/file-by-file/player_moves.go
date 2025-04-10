// Package tic_tac_toe handles player moves within the Tic Tac Toe game.
package tic_tac_toe

// Board represents the Tic-Tac-Toe board.
type Board [][]string

// isValidMove checks if a move is valid given the current board state.
func isValidMove(board Board, x int, y int) bool {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return false
	}
	return board[x][y] == " "
}

// makeMove updates the board with a player's move if it's valid.
func makeMove(board Board, x int, y int, player string) bool {
	if isValidMove(board, x, y) {
		board[x][y] = player
		return true
	}
	return false
}

// NewBoard initializes an empty Tic-Tac-Toe board.
func NewBoard() Board {
	return Board{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}
