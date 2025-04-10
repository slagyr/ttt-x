package tic_tac_toe

import (
	"math/rand"
	"time"
)

// Board represents the Tic-Tac-Toe board.
type Board [][]string

// GetBotMove calculates and returns the bot's move with improved strategy.
func GetBotMove(board Board, botPlayer string, opponent string) (int, int) {
	// Initialize the random seed for random choices
	rand.Seed(time.Now().UnixNano())

	// Check if the bot can win in the next move
	for i, row := range board {
		for j, cell := range row {
			if cell == " " {
				// Assume the bot moves here
				board[i][j] = botPlayer
				if CheckWin(board, botPlayer) {
					board[i][j] = " " // reset
					return i, j
				}
				board[i][j] = " " // reset
			}
		}
	}

	// Check if the opponent can win in the next move and block
	for i, row := range board {
		for j, cell := range row {
			if cell == " " {
				// Assume the opponent moves here
				board[i][j] = opponent
				if CheckWin(board, opponent) {
					board[i][j] = " " // reset
					return i, j
				}
				board[i][j] = " " // reset
			}
		}
	}

	// Choose the center if available
	if board[1][1] == " " {
		return 1, 1
	}

	// Choose a corner if available
	corners := []struct{ x, y int }{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
	availableCorners := make([]struct{ x, y int }, 0)
	for _, c := range corners {
		if board[c.x][c.y] == " " {
			availableCorners = append(availableCorners, c)
		}
	}
	if len(availableCorners) > 0 {
		choice := availableCorners[rand.Intn(len(availableCorners))]
		return choice.x, choice.y
	}

	// Else choose any available space at random
	emptyPositions := []struct{ x, y int }{}
	for i, row := range board {
		for j, cell := range row {
			if cell == " " {
				emptyPositions = append(emptyPositions, struct{ x, y int }{i, j})
			}
		}
	}

	if len(emptyPositions) > 0 {
		choice := emptyPositions[rand.Intn(len(emptyPositions))]
		return choice.x, choice.y
	}

	// No available moves
	return -1, -1
}

// CheckWin evaluates whether the given player has won the game.
func CheckWin(board Board, player string) bool {
	// Implement a win-check for the board
	// This function should check rows, columns, and diagonals for a win condition
	// Code for this function is similar to the implementation in previous messages
	return false // Placeholder return for compilation, replace with actual check
}
