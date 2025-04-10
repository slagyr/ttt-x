package board

import (
	"testing"
)

func TestEvaluateBoard(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		expected interface{}
	}{
		{
			name: "O wins (row)",
			board: Board{
				"O", "O", "O",
				"X", "X", "O",
				"X", "8", "X",
			},
			expected: 10,
		},
		{
			name: "O wins (diagonal)",
			board: Board{
				"O", "2", "X",
				"O", "X", "X",
				"O", "X", "O",
			},
			expected: 10,
		},
		{
			name: "X wins (diagonal)",
			board: Board{
				"1", "2", "X",
				"4", "X", "O",
				"X", "7", "O",
			},
			expected: -10,
		},
		{
			name: "Tie",
			board: Board{
				"X", "X", "O",
				"O", "O", "X",
				"X", "O", "X",
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvaluateBoard(tt.board)
			if got != tt.expected {
				t.Errorf("EvaluateBoard() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetPaths(t *testing.T) {
	board := NewBoard()
	paths := GetPaths(board)
	
	// Should have 8 paths (3 rows, 3 columns, 2 diagonals)
	if len(paths) != 8 {
		t.Errorf("Expected 8 paths, got %d", len(paths))
	}
	
	// Each path should have 3 cells
	for i, path := range paths {
		if len(path) != 3 {
			t.Errorf("Path %d: expected 3 cells, got %d", i, len(path))
		}
	}
	
	// Test with a specific board
	board = Board{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	paths = GetPaths(board)
	
	expectedPaths := [][]Cell{
		{"1", "2", "3"}, // top row
		{"4", "5", "6"}, // middle row
		{"7", "8", "9"}, // bottom row
		{"1", "4", "7"}, // left column
		{"2", "5", "8"}, // middle column
		{"3", "6", "9"}, // right column
		{"1", "5", "9"}, // diagonal top-left to bottom-right
		{"3", "5", "7"}, // diagonal top-right to bottom-left
	}
	
	for i, expectedPath := range expectedPaths {
		for j, cell := range expectedPath {
			if paths[i][j] != cell {
				t.Errorf("Path %d, cell %d: expected %s, got %s", i, j, cell, paths[i][j])
			}
		}
	}
	
	// Test with a board that has some moves played
	board = Board{
		"O", "O", "O",
		"4", "X", "6",
		"7", "X", "X",
	}
	paths = GetPaths(board)
	
	expectedPaths = [][]Cell{
		{"O", "O", "O"}, // top row
		{"4", "X", "6"}, // middle row
		{"7", "X", "X"}, // bottom row
		{"O", "4", "7"}, // left column
		{"O", "X", "X"}, // middle column
		{"O", "6", "X"}, // right column
		{"O", "X", "X"}, // diagonal top-left to bottom-right
		{"O", "X", "7"}, // diagonal top-right to bottom-left
	}
	
	for i, expectedPath := range expectedPaths {
		for j, cell := range expectedPath {
			if paths[i][j] != cell {
				t.Errorf("Path %d, cell %d: expected %s, got %s", i, j, cell, paths[i][j])
			}
		}
	}
}

func TestScore(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		expected string
	}{
		{
			name:     "Empty board is in progress",
			board:    Board{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			expected: "in-progress",
		},
		{
			name: "Partially filled board with no winner is in progress",
			board: Board{
				"X", "O", "X",
				"O", "O", "6",
				"7", "8", "9",
			},
			expected: "in-progress",
		},
		{
			name: "X wins horizontally",
			board: Board{
				"X", "X", "X",
				"4", "5", "6",
				"7", "8", "9",
			},
			expected: "X wins!",
		},
		{
			name: "X wins vertically",
			board: Board{
				"X", "2", "3",
				"X", "5", "6",
				"X", "8", "9",
			},
			expected: "X wins!",
		},
		{
			name: "O wins horizontally",
			board: Board{
				"1", "2", "3",
				"O", "O", "O",
				"7", "8", "9",
			},
			expected: "O wins!",
		},
		{
			name: "Tie",
			board: Board{
				"X", "O", "X",
				"X", "X", "O",
				"O", "X", "O",
			},
			expected: "It's a tie!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Score(tt.board)
			if got != tt.expected {
				t.Errorf("Score() = %v, want %v", got, tt.expected)
			}
		})
	}
}