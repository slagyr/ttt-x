package board

import (
	"reflect"
	"testing"
)

// TestEvaluateBoard tests the evaluateBoard function
func TestEvaluateBoard(t *testing.T) {
	tests := []struct {
		name     string
		board    []interface{}
		expected int
	}{
		{
			"O wins horizontally",
			[]interface{}{"o", "o", "o", "x", "x", "o", "x", 8, "x"},
			10,
		},
		{
			"O wins diagonally",
			[]interface{}{"o", 2, "x", "o", "x", "x", "o", "x", "o"},
			10,
		},
		{
			"X wins vertically",
			[]interface{}{1, 2, "x", 4, "x", "o", "x", 7, "o"},
			-10,
		},
		{
			"Tie game",
			[]interface{}{"x", "x", "o", "o", "o", "x", "x", "o", "x"},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evaluateBoard(tt.board)
			if result != tt.expected {
				t.Errorf("evaluateBoard() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestToPaths tests the separation of board into winning paths
func TestToPaths(t *testing.T) {
	tests := []struct {
		name     string
		board    []interface{}
		expected [][]interface{}
	}{
		{
			"Empty board",
			make([]interface{}, 9),
			[][]interface{}{
				{"", "", ""}, {"", "", ""}, {"", "", ""},
				{"", "", ""}, {"", "", ""}, {"", "", ""},
				{"", "", ""}, {"", "", ""},
			},
		},
		{
			"Numbered board",
			[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[][]interface{}{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
				{1, 4, 7}, {2, 5, 8}, {3, 6, 9},
				{1, 5, 9}, {3, 5, 7},
			},
		},
		{
			"Mixed board",
			[]interface{}{"o", "o", "o", 4, "x", 6, 7, "x", "x"},
			[][]interface{}{
				{"o", "o", "o"}, {4, "x", 6}, {7, "x", "x"},
				{"o", 4, 7}, {"o", "x", "x"}, {"o", 6, "x"},
				{"o", "x", "x"}, {"o", "x", 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toPaths(tt.board)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("toPaths() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestScore tests the scoring of the game
func TestScore(t *testing.T) {
	tests := []struct {
		name     string
		board    []interface{}
		expected string
	}{
		{
			"Empty board in progress",
			make([]interface{}, 9),
			"in-progress",
		},
		{
			"Partial board in progress",
			[]interface{}{"x", "o", "x", "o", "o", 6, 7, 8, 9},
			"in-progress",
		},
		{
			"Numbered board in progress",
			[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
			"in-progress",
		},
		{
			"X wins top row",
			[]interface{}{"x", "x", "x", 4, 5, 6, 7, 8, 9},
			"X wins!",
		},
		{
			"X wins middle row",
			[]interface{}{1, 2, 3, "x", "x", "x", 7, 8, 9},
			"X wins!",
		},
		{
			"O wins middle row",
			[]interface{}{1, 2, 3, "o", "o", "o", 7, 8, 9},
			"O wins!",
		},
		{
			"Tie game",
			[]interface{}{"x", "o", "x", "x", "x", "o", "o", "x", "o"},
			"It's a tie!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := score(tt.board)
			if result != tt.expected {
				t.Errorf("score() = %v, want %v", result, tt.expected)
			}
		})
	}
}
