package bot

import (
	"tictactoe/board"
	"tictactoe/utils"
)

// PlayBotTurn plays a turn for the bot and returns the updated board
func PlayBotTurn(board []interface{}, player string) []interface{} {
	move := FindBestMove(board, player)
	utils.DisplayBotMoveMessage(move)
	
	newBoard := make([]interface{}, len(board))
	copy(newBoard, board)
	newBoard[move-1] = player
	
	return newBoard
}

// FindBestMove finds the best move for the bot using minimax algorithm
func FindBestMove(board []interface{}, player string) int {
	var bestMove int
	bestScore := -1000
	
	// Try each available move
	for _, move := range getAvailableMoves(board) {
		// Create a new board with the move
		newBoard := make([]interface{}, len(board))
		copy(newBoard, board)
		newBoard[move-1] = player
		
		// Calculate score for this move
		score := minimax(newBoard, false, 0)
		
		// Update best move if we found a better score
		if scoreValue(score) > bestScore {
			bestScore = scoreValue(score)
			bestMove = move
		}
	}
	
	return bestMove
}

// Helper function to convert interface{} to int for score comparison
func scoreValue(score interface{}) int {
	if val, ok := score.(int); ok {
		return val
	}
	return 0
}

// minimax implements the minimax algorithm for finding optimal moves
func minimax(b []interface{}, isMax bool, depth int) interface{} {
	// Check game state
	result := board.EvaluateBoard(b)
	
	// If game is over, return the score
	if result != "in-progress" {
		return result
	}
	
	// If it's maximizer's turn (bot)
	if isMax {
		return maxMove(b, board.O, depth)
	}
	
	// If it's minimizer's turn (player)
	return minMove(b, board.X, depth)
}

// maxMove finds the best move for maximizer (bot/O)
func maxMove(b []interface{}, player string, depth int) int {
	bestScore := -1000
	
	// Try each available move
	for _, move := range getAvailableMoves(b) {
		// Create a new board with the move
		newBoard := make([]interface{}, len(b))
		copy(newBoard, b)
		newBoard[move-1] = player
		
		// Calculate score for this move
		score := minimax(newBoard, false, depth+1)
		
		// Update best score
		scoreInt := scoreValue(score)
		if scoreInt > bestScore {
			bestScore = scoreInt
		}
	}
	
	return bestScore
}

// minMove finds the best move for minimizer (player/X)
func minMove(b []interface{}, player string, depth int) int {
	bestScore := 1000
	
	// Try each available move
	for _, move := range getAvailableMoves(b) {
		// Create a new board with the move
		newBoard := make([]interface{}, len(b))
		copy(newBoard, b)
		newBoard[move-1] = player
		
		// Calculate score for this move
		score := minimax(newBoard, true, depth+1)
		
		// Update best score
		scoreInt := scoreValue(score)
		if scoreInt < bestScore {
			bestScore = scoreInt
		}
	}
	
	return bestScore
}

// getAvailableMoves returns a list of available moves on the board
func getAvailableMoves(board []interface{}) []int {
	var moves []int
	
	for _, cell := range board {
		if val, ok := cell.(int); ok {
			moves = append(moves, val)
		}
	}
	
	return moves
}