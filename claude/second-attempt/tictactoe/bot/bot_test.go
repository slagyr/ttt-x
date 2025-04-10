package bot

import (
	"testing"
	"tictactoe/board"
)

func TestFindBestMove(t *testing.T) {
	tests := []struct {
		name     string
		board    []interface{}
		player   string
		expected int
	}{
		{
			name:     "Find winning move (row)",
			board:    []interface{}{board.O, 2, board.O, board.X, board.X, board.O, board.X, 8, board.X},
			player:   board.O,
			expected: 2,
		},
		{
			name:     "Find winning move (another row)",
			board:    []interface{}{board.O, 2, board.O, board.X, board.X, 6, board.X, 8, 9},
			player:   board.O,
			expected: 2,
		},
		{
			name:     "Find winning move (diagonal)",
			board:    []interface{}{1, board.X, board.X, 4, board.O, board.O, board.O, board.X, board.X},
			player:   board.O,
			expected: 4,
		},
		{
			name:     "Find winning move (another diagonal)",
			board:    []interface{}{1, 2, 3, board.X, board.X, board.O, 7, board.X, board.O},
			player:   board.O,
			expected: 3,
		},
		{
			name:     "Find blocking move (row)",
			board:    []interface{}{1, 2, board.O, 4, board.X, board.X, 7, 8, 9},
			player:   board.O,
			expected: 4,
		},
		{
			name:     "Find blocking move (column)",
			board:    []interface{}{1, 2, board.O, 4, board.X, 6, 7, 8, board.X},
			player:   board.O,
			expected: 1,
		},
		{
			name:     "Find blocking move (diagonal)",
			board:    []interface{}{board.X, board.O, board.X, board.O, board.O, board.X, 7, board.X, 9},
			player:   board.O,
			expected: 9,
		},
		{
			name:     "Find best strategic move",
			board:    []interface{}{board.X, board.O, board.X, 4, board.O, 6, board.X, board.X, board.O},
			player:   board.O,
			expected: 4,
		},
		{
			name:     "Find best potential move",
			board:    []interface{}{board.O, board.X, board.O, board.X, 5, 6, 7, 8, board.X},
			player:   board.O,
			expected: 5,
		},
		{
			name:     "Play center if first move is a corner (top-left)",
			board:    []interface{}{board.X, 2, 3, 4, 5, 6, 7, 8, 9},
			player:   board.O,
			expected: 5,
		},
		{
			name:     "Play center if first move is a corner (top-right)",
			board:    []interface{}{1, 2, board.X, 4, 5, 6, 7, 8, 9},
			player:   board.O,
			expected: 5,
		},
		{
			name:     "Play center if first move is a corner (bottom-left)",
			board:    []interface{}{1, 2, 3, 4, 5, 6, board.X, 8, 9},
			player:   board.O,
			expected: 5,
		},
		{
			name:     "Play center if first move is a corner (bottom-right)",
			board:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8, board.X},
			player:   board.O,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindBestMove(tt.board, tt.player)
			if result != tt.expected {
				t.Errorf("FindBestMove() = %v, expected %v", result, tt.expected)
			}
		})
	}
}