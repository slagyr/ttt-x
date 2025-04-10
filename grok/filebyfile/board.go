package board

// toPaths converts a board into all possible winning paths
// Returns top, middle, bottom rows, left, middle, right columns, and both diagonals
func toPaths(board []interface{}) [][]interface{} {
	paths := make([][]interface{}, 0, 8)

	// Horizontal rows (top, middle, bottom)
	for i := 0; i < 9; i += 3 {
		paths = append(paths, board[i:i+3])
	}

	// Vertical columns (left, middle, right)
	for i := 0; i < 3; i++ {
		column := make([]interface{}, 3)
		for j := 0; j < 3; j++ {
			column[j] = board[j*3+i]
		}
		paths = append(paths, column)
	}

	// Diagonals
	// Left-to-right (0,4,8)
	diagonal1 := []interface{}{board[0], board[4], board[8]}
	paths = append(paths, diagonal1)

	// Right-to-left (2,4,6)
	diagonal2 := []interface{}{board[2], board[4], board[6]}
	paths = append(paths, diagonal2)

	return paths
}

// score determines the current state of the game
func score(board []interface{}) string {
	paths := toPaths(board)

	// Check for X win
	for _, path := range paths {
		if allEqual(path, "x") {
			return "X wins!"
		}
	}

	// Check for O win
	for _, path := range paths {
		if allEqual(path, "o") {
			return "O wins!"
		}
	}

	// Check for tie (board full)
	if allFilled(board) {
		return "It's a tie!"
	}

	return "in-progress"
}

// evaluateBoard assigns a numeric score to the board state
func evaluateBoard(board []interface{}) int {
	paths := toPaths(board)

	// Check for X win
	for _, path := range paths {
		if allEqual(path, "x") {
			return -10
		}
	}

	// Check for O win
	for _, path := range paths {
		if allEqual(path, "o") {
			return 10
		}
	}

	// Check for tie (board full)
	if allFilled(board) {
		return 0
	}

	return 0 // in-progress
}

// allEqual checks if all elements in a slice equal the given value
func allEqual(path []interface{}, value string) bool {
	for _, v := range path {
		if v != value {
			return false
		}
	}
	return true
}

// allFilled checks if the board has no empty spaces
func allFilled(board []interface{}) bool {
	for _, v := range board {
		if v != "x" && v != "o" {
			return false
		}
	}
	return true
}
