package board

// Interface definition not needed in Go version

const (
	// X represents the X player mark
	X string = "X"
	// O represents the O player mark
	O string = "O"
)

// GetPaths returns all possible winning paths on the board
// Top, middle, bottom, left, middle, right, l-r, r-l
func GetPaths(board []interface{}) [][]interface{} {
	paths := make([][]interface{}, 8)
	
	// Rows
	paths[0] = []interface{}{board[0], board[1], board[2]}
	paths[1] = []interface{}{board[3], board[4], board[5]}
	paths[2] = []interface{}{board[6], board[7], board[8]}
	
	// Columns
	paths[3] = []interface{}{board[0], board[3], board[6]}
	paths[4] = []interface{}{board[1], board[4], board[7]}
	paths[5] = []interface{}{board[2], board[5], board[8]}
	
	// Diagonals
	paths[6] = []interface{}{board[0], board[4], board[8]}
	paths[7] = []interface{}{board[2], board[4], board[6]}
	
	return paths
}

// Score returns the game state as a string: "X wins!", "O wins!", "It's a tie!", or "in-progress"
func Score(board []interface{}) string {
	paths := GetPaths(board)
	
	// Check for X win
	for _, path := range paths {
		if allEqual(path, X) {
			return "X wins!"
		}
	}
	
	// Check for O win
	for _, path := range paths {
		if allEqual(path, O) {
			return "O wins!"
		}
	}
	
	// Check for tie
	if isBoardFull(board) {
		return "It's a tie!"
	}
	
	return "in-progress"
}

// EvaluateBoard returns the numeric score of the board for minimax:
// 10 for O win, -10 for X win, 0 for tie, "in-progress" string for ongoing game
func EvaluateBoard(board []interface{}) interface{} {
	paths := GetPaths(board)
	
	// Check for X win
	for _, path := range paths {
		if allEqual(path, X) {
			return -10
		}
	}
	
	// Check for O win
	for _, path := range paths {
		if allEqual(path, O) {
			return 10
		}
	}
	
	// Check for tie
	if isBoardFull(board) {
		return 0
	}
	
	return "in-progress"
}

// Helper function to check if all elements in a path are equal to the specified player
func allEqual(path []interface{}, player string) bool {
	for _, mark := range path {
		if mark != player {
			return false
		}
	}
	return true
}

// Helper function to check if the board is full
func isBoardFull(board []interface{}) bool {
	for _, cell := range board {
		_, isNumber := cell.(int)
		// Empty string cells are considered available
		_, isEmptyString := cell.(string)
		if isNumber || (isEmptyString && cell.(string) == "") {
			return false
		}
	}
	return true
}