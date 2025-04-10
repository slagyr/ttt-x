package utils

import (
	"fmt"
)

// DisplayInvalidMoveError displays an error message for invalid moves
func DisplayInvalidMoveError() {
	fmt.Println("Please choose an empty cell")
}

// PromptUserForMove prompts the user to enter a move
func PromptUserForMove() {
	fmt.Println("Please enter your move (type 1-9 and hit enter):")
}

// DisplayBotMoveMessage displays a message about the bot's move
func DisplayBotMoveMessage(move int) {
	fmt.Printf("Your opponent plays %d!\n", move)
}

// PrintBoard prints the current state of the board
func PrintBoard(board []interface{}) {
	formattedBoard := formatBoard(board)
	
	fmt.Println(formattedBoard[0], formattedBoard[1], formattedBoard[2])
	fmt.Println(formattedBoard[3], formattedBoard[4], formattedBoard[5])
	fmt.Println(formattedBoard[6], formattedBoard[7], formattedBoard[8])
}

// formatBoard formats the board cells for display
func formatBoard(board []interface{}) []string {
	result := make([]string, len(board))
	
	purple := "\u001b[35m"
	reset := "\u001b[0m"
	
	for i, cell := range board {
		switch v := cell.(type) {
		case string:
			// Handle player marks (X and O)
			result[i] = v
		case int:
			result[i] = fmt.Sprintf("%s%d%s", purple, v, reset)
		default:
			result[i] = fmt.Sprintf("%v", cell)
		}
	}
	
	return result
}