package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ttt/board"
)

// GetHumanMove prompts the user for a valid move.
func GetHumanMove(b *board.Board) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Choose position to place your 'X' (0-8): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		move, err := strconv.Atoi(input)
		if err == nil && move >= 0 && move <= 8 && b.IsPositionEmpty(move) {
			return move
		}
		fmt.Println("Invalid position, please try again.")
	}
}

// GetComputerMove uses minimax to choose the best move for "O".
func GetComputerMove(b *board.Board) int {
	_, move := minimax(b, "O")
	return move
}

// minimax implements the minimax algorithm, returning score and best move.
func minimax(b *board.Board, player string) (int, int) {
	winner := b.GetWinner()
	if winner == "O" {
		return 10, -1
	}
	if winner == "X" {
		return -10, -1
	}
	if b.IsFull() {
		return 0, -1
	}

	if player == "O" {
		bestScore := -1000
		bestMove := -1
		for _, pos := range b.AvailablePositions() {
			newBoard := b.Copy()
			newBoard.Mark(pos, "O")
			score, _ := minimax(newBoard, "X")
			if score > bestScore {
				bestScore = score
				bestMove = pos
			}
		}
		return bestScore, bestMove
	} else { // player == "X"
		bestScore := 1000
		bestMove := -1
		for _, pos := range b.AvailablePositions() {
			newBoard := b.Copy()
			newBoard.Mark(pos, "X")
			score, _ := minimax(newBoard, "O")
			if score < bestScore {
				bestScore = score
				bestMove = pos
			}
		}
		return bestScore, bestMove
	}
}
