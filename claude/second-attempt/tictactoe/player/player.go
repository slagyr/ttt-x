package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tictactoe/utils"
)

// PlayUserTurn plays a turn for the user and returns the updated board
func PlayUserTurn(board []interface{}, player string) []interface{} {
	utils.PromptUserForMove()
	
	for {
		move := readMove()
		
		// Check if move is valid
		if move >= 0 && move < len(board) {
			if _, isNumber := board[move].(int); isNumber {
				// Make a copy of the board to avoid modifying the original
				newBoard := make([]interface{}, len(board))
				copy(newBoard, board)
				newBoard[move] = player
				
				return newBoard
			}
		}
		
		utils.DisplayInvalidMoveError()
	}
}

// readMove reads a move from standard input
// Returns the index (0-8) of the chosen cell
func readMove() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return -1 // Error reading input
	}
	
	// Clean input (remove whitespace, newlines)
	input = strings.TrimSpace(input)
	
	// Try to parse the input as an integer
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return -1 // Invalid input
	}
	
	// Convert from 1-9 notation to 0-8 array index
	return num - 1
}