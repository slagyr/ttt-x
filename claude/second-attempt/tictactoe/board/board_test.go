package board

import (
	"reflect"
	"testing"
)

func TestGetPaths(t *testing.T) {
	tests := []struct {
		name     string
		board    []interface{}
		expected [][]interface{}
	}{
		{
			name:  "Empty board paths",
			board: []interface{}{"", "", "", "", "", "", "", "", ""},
			expected: [][]interface{}{
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
			},
		},
		{
			name:  "Numbered board paths",
			board: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: [][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
				{1, 5, 9},
				{3, 5, 7},
			},
		},
		{
			name:  "Mixed board paths",
			board: []interface{}{O, O, O, 4, X, 6, 7, X, X},
			expected: [][]interface{}{
				{O, O, O},
				{4, X, 6},
				{7, X, X},
				{O, 4, 7},
				{O, X, X},
				{O, 6, X},
				{O, X, X},
				{O, X, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetPaths(tt.board)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GetPaths() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestEvaluateBoard(t *testing.T) {
	tests := []struct {
		name     string
		board    []interface{}
		expected interface{}
	}{
		{
			name:     "O wins (row)",
			board:    []interface{}{O, O, O, X, X, O, X, 8, X},
			expected: 10,
		},
		{
			name:     "O wins (column)",
			board:    []interface{}{O, 2, X, O, X, X, O, X, O},
			expected: 10,
		},
		{
			name:     "X wins (diagonal)",
			board:    []interface{}{1, 2, X, 4, X, O, X, 7, O},
			expected: -10,
		},
		{
			name:     "Tie",
			board:    []interface{}{X, X, O, O, O, X, X, O, X},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateBoard(tt.board)
			if result != tt.expected {
				t.Errorf("EvaluateBoard() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestScore(t *testing.T) {
	tests := []struct {
		name     string
		board    []interface{}
		expected string
	}{
		{
			name:     "Empty board is in progress",
			board:    []interface{}{"", "", "", "", "", "", "", "", ""},
			expected: "in-progress",
		},
		{
			name:     "Partially filled board is in progress",
			board:    []interface{}{X, O, X, O, O, 6, 7, 8, 9},
			expected: "in-progress",
		},
		{
			name:     "Numbered board is in progress",
			board:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: "in-progress",
		},
		{
			name:     "X wins top row",
			board:    []interface{}{X, X, X, 4, 5, 6, 7, 8, 9},
			expected: "X wins!",
		},
		{
			name:     "X wins middle row",
			board:    []interface{}{1, 2, 3, X, X, X, 7, 8, 9},
			expected: "X wins!",
		},
		{
			name:     "O wins middle row",
			board:    []interface{}{1, 2, 3, O, O, O, 7, 8, 9},
			expected: "O wins!",
		},
		{
			name:     "Tie game",
			board:    []interface{}{X, O, X, X, X, O, O, X, O},
			expected: "It's a tie!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Score(tt.board)
			if result != tt.expected {
				t.Errorf("Score() = %v, expected %v", result, tt.expected)
			}
		})
	}
}