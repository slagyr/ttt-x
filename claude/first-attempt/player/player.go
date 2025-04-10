package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tic-tac-toe/board"
)

// PlayUserTurn handles user input and updates the board
func PlayUserTurn(b board.Board, player string) board.Board {
	fmt.Println("Please enter your move (type 1-9 and hit enter):")
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		move, err := strconv.Atoi(input)
		if err != nil || move < 1 || move > 9 {
			fmt.Println("Please enter a number between 1 and 9")
			continue
		}
		
		// Check if the move is valid (cell contains a number)
		idx := move - 1
		if !board.IsNumber(b[idx]) {
			fmt.Println("Please choose an empty cell")
			continue
		}
		
		// Update the board
		b[idx] = board.Cell(player)
		return b
	}
}

// PlayBotTurn uses minimax to find and play the best move
func PlayBotTurn(b board.Board, player string) board.Board {
	move := FindBestMove(b, player)
	fmt.Printf("Your opponent plays %d!\n", move+1)
	
	// Update the board
	b[move] = board.Cell(player)
	return b
}

// FindBestMove uses minimax algorithm to find the best move
func FindBestMove(b board.Board, player string) int {
	bestScore := -1000
	bestMove := -1
	
	// Try each available move
	for i := 0; i < 9; i++ {
		if board.IsNumber(b[i]) {
			// Try this move
			newBoard := b
			newBoard[i] = board.Cell(player)
			
			// Evaluate move with minimax
			score := Minimax(newBoard, false, 0)
			
			// If this move is better than our best move so far, update bestMove
			if v, ok := score.(int); ok && v > bestScore {
				bestScore = v
				bestMove = i
			}
		}
	}
	
	return bestMove
}

// Minimax is the recursive minimax algorithm
func Minimax(b board.Board, isMaximizing bool, depth int) interface{} {
	// Check if the game is over
	result := board.EvaluateBoard(b)
	if result != "in-progress" {
		return result
	}
	
	if isMaximizing {
		return MaxMove(b, "O", depth)
	}
	return MinMove(b, "X", depth)
}

// MaxMove implements the maximizing part of minimax
func MaxMove(b board.Board, player string, depth int) int {
	bestScore := -1000
	
	// Try each available move
	for i := 0; i < 9; i++ {
		if board.IsNumber(b[i]) {
			// Try this move
			newBoard := b
			newBoard[i] = board.Cell(player)
			
			// Evaluate move with minimax
			score := Minimax(newBoard, false, depth+1)
			
			// Update best score
			if v, ok := score.(int); ok {
				bestScore = max(bestScore, v)
			}
		}
	}
	
	return bestScore
}

// MinMove implements the minimizing part of minimax
func MinMove(b board.Board, player string, depth int) int {
	bestScore := 1000
	
	// Try each available move
	for i := 0; i < 9; i++ {
		if board.IsNumber(b[i]) {
			// Try this move
			newBoard := b
			newBoard[i] = board.Cell(player)
			
			// Evaluate move with minimax
			score := Minimax(newBoard, true, depth+1)
			
			// Update best score
			if v, ok := score.(int); ok {
				bestScore = min(bestScore, v)
			}
		}
	}
	
	return bestScore
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}