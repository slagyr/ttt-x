package main

func GetBestMove(board *Board, player string) (int, int) {
	bestScore := -1000
	var bestMove [2]int
	// Go through all cells and choose the best move using Minimax
	for _, move := range board.AvailableMoves() {
		board.MakeMove(player, move[0], move[1])
		score := minimax(board, 0, false, player)
		board.MakeMove("", move[0], move[1]) // Undo move
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove[0], bestMove[1]
}

func minimax(board *Board, depth int, isMaximizing bool, computer string) int {
	human := "X"
	if computer == "X" {
		human = "O"
	}

	winner := board.CheckWinner()
	if winner == computer {
		return 10 - depth
	} else if winner == human {
		return depth - 10
	} else if board.IsFull() {
		return 0
	}

	if isMaximizing {
		bestScore := -1000
		for _, move := range board.AvailableMoves() {
			board.MakeMove(computer, move[0], move[1])
			score := minimax(board, depth+1, false, computer)
			board.MakeMove("", move[0], move[1])
			bestScore = max(score, bestScore)
		}
		return bestScore
	} else {
		bestScore := 1000
		for _, move := range board.AvailableMoves() {
			board.MakeMove(human, move[0], move[1])
			score := minimax(board, depth+1, true, computer)
			board.MakeMove("", move[0], move[1])
			bestScore = min(score, bestScore)
		}
		return bestScore
	}
}

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
