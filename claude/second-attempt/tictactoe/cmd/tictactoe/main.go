package main

import (
	"tictactoe/board"
	"tictactoe/bot"
	"tictactoe/player"
	"tictactoe/utils"
)

func main() {
	// Initialize board with numbers 1-9
	initialBoard := make([]interface{}, 9)
	for i := 0; i < 9; i++ {
		initialBoard[i] = i + 1
	}
	
	// Player sequence (alternating X and O)
	currentPlayer := board.X
	
	// Game loop
	gameBoard := initialBoard
	
	for {
		utils.PrintBoard(gameBoard)
		
		// Check if game is over
		score := board.Score(gameBoard)
		if score != "in-progress" {
			println(score)
			break
		}
		
		// Play turn based on current player
		if currentPlayer == board.X {
			gameBoard = player.PlayUserTurn(gameBoard, currentPlayer)
		} else {
			gameBoard = bot.PlayBotTurn(gameBoard, currentPlayer)
		}
		
		// Switch player
		if currentPlayer == board.X {
			currentPlayer = board.O
		} else {
			currentPlayer = board.X
		}
	}
}