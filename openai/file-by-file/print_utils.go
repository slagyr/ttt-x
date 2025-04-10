// Package tic_tac_toe provides utility functions for printing within the Tic Tac Toe game.
package tic_tac_toe

import (
	"fmt"
)

// printBoard is a function to print the Tic-Tac-Toe board in the console.
func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
}

// printInstructions prints the game instructions to the console.
func printInstructions() {
	fmt.Println("Welcome to Tic-Tac-Toe!")
	fmt.Println("Here are the instructions...")
	// Add more instruction details here
}

// Additional helper functions would go here
// func exampleHelperFunction(param int) string {
//     // Implementation
//     return ""
// }
