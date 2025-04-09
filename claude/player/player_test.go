package player

import (
	"testing"
	"tic-tac-toe/board"
)

func TestFindBestMove(t *testing.T) {
	tests := []struct {
		name          string
		board         board.Board
		player        string
		expectedMove  int
		testCategory  string
	}{
		// Test cases for finding winning moves
		{
			name: "Finds winning move (row)",
			board: board.Board{
				"O", "2", "O",
				"X", "X", "O",
				"X", "8", "X",
			},
			player:       "O",
			expectedMove: 1, // index 1 is the 2nd position (0-indexed)
			testCategory: "winning move",
		},
		{
			name: "Finds winning move (another row)",
			board: board.Board{
				"O", "2", "O",
				"X", "X", "6",
				"X", "8", "9",
			},
			player:       "O",
			expectedMove: 1, // index 1 is the 2nd position
			testCategory: "winning move",
		},
		{
			name: "Finds winning move (column)",
			board: board.Board{
				"1", "X", "X",
				"4", "O", "O",
				"O", "X", "X",
			},
			player:       "O",
			expectedMove: 3, // index 3 is the 4th position
			testCategory: "winning move",
		},
		{
			name: "Finds winning move (diagonal)",
			board: board.Board{
				"1", "2", "3",
				"X", "X", "O",
				"7", "X", "O",
			},
			player:       "O",
			expectedMove: 2, // index 2 is the 3rd position
			testCategory: "winning move",
		},

		// Test cases for blocking opponent's winning move
		{
			name: "Blocks opponent's winning move (row)",
			board: board.Board{
				"1", "2", "O",
				"4", "X", "X",
				"7", "8", "9",
			},
			player:       "O",
			expectedMove: 3, // index 3 is the 4th position
			testCategory: "blocking move",
		},
		{
			name: "Blocks opponent's winning move (diagonal)",
			board: board.Board{
				"1", "2", "O",
				"4", "X", "6",
				"7", "8", "X",
			},
			player:       "O",
			expectedMove: 0, // index 0 is the 1st position
			testCategory: "blocking move",
		},
		{
			name: "Blocks opponent's winning move (complex board)",
			board: board.Board{
				"X", "O", "X",
				"O", "O", "X",
				"7", "X", "O",
			},
			player:       "O",
			expectedMove: 6, // index 6 is the 7th position
			testCategory: "blocking move",
		},
		{
			name: "Blocks opponent's winning move (another complex board)",
			board: board.Board{
				"X", "O", "X",
				"4", "O", "6",
				"X", "X", "O",
			},
			player:       "O",
			expectedMove: 3, // index 3 is the 4th position
			testCategory: "blocking move",
		},

		// Strategic moves
		{
			name: "Plays center if first move is a corner",
			board: board.Board{
				"X", "2", "3",
				"4", "5", "6",
				"7", "8", "9",
			},
			player:       "O",
			expectedMove: 4, // index 4 is the center position
			testCategory: "strategic move",
		},
		{
			name: "Plays center if opponent plays a corner",
			board: board.Board{
				"1", "2", "X",
				"4", "5", "6",
				"7", "8", "9",
			},
			player:       "O",
			expectedMove: 4, // index 4 is the center position
			testCategory: "strategic move",
		},
		{
			name: "Plays center if opponent plays bottom left corner",
			board: board.Board{
				"1", "2", "3",
				"4", "5", "6",
				"X", "8", "9",
			},
			player:       "O",
			expectedMove: 4, // index 4 is the center position
			testCategory: "strategic move",
		},
		{
			name: "Plays center if opponent plays bottom right corner",
			board: board.Board{
				"1", "2", "3",
				"4", "5", "6",
				"7", "8", "X",
			},
			player:       "O",
			expectedMove: 4, // index 4 is the center position
			testCategory: "strategic move",
		},
		{
			name: "Chooses strategic position with most winning potential",
			board: board.Board{
				"O", "X", "O",
				"X", "5", "6",
				"7", "8", "X",
			},
			player:       "O",
			expectedMove: 4, // index 4 is the center position
			testCategory: "strategic move",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			move := FindBestMove(tt.board, tt.player)
			if move != tt.expectedMove {
				t.Errorf("\nTest category: %s\nFindBestMove(%v, %s) = %d, expected %d", 
					tt.testCategory, tt.board, tt.player, move, tt.expectedMove)
			}
		})
	}
}

func TestMinimaxFunctions(t *testing.T) {
	// Test Minimax function returns expected values
	t.Run("Minimax returns correct values for terminal states", func(t *testing.T) {
		// X has won - should return -10
		xWinsBoard := board.Board{
			"X", "X", "X",
			"O", "O", "6",
			"7", "8", "9",
		}
		result := Minimax(xWinsBoard, true, 0)
		if result != -10 {
			t.Errorf("Minimax with X winning should return -10, got %v", result)
		}
		
		// O has won - should return 10
		oWinsBoard := board.Board{
			"X", "X", "O",
			"X", "O", "6",
			"O", "8", "9",
		}
		result = Minimax(oWinsBoard, false, 0)
		if result != 10 {
			t.Errorf("Minimax with O winning should return 10, got %v", result)
		}
		
		// Tie - should return 0
		tieBoard := board.Board{
			"X", "O", "X",
			"X", "O", "X",
			"O", "X", "O",
		}
		result = Minimax(tieBoard, true, 0)
		if result != 0 {
			t.Errorf("Minimax with tie should return 0, got %v", result)
		}
	})
	
	// Test MaxMove function
	t.Run("MaxMove returns best score for maximizing player", func(t *testing.T) {
		// One move away from win for O
		board := board.Board{
			"O", "O", "3",
			"X", "X", "6",
			"7", "8", "9",
		}
		score := MaxMove(board, "O", 0)
		if score != 10 {
			t.Errorf("MaxMove should choose winning move with score 10, got %d", score)
		}
	})
	
	// Test MinMove function
	t.Run("MinMove returns best score for minimizing player", func(t *testing.T) {
		// One move away from win for X
		board := board.Board{
			"X", "X", "3",
			"O", "O", "6",
			"7", "8", "9",
		}
		score := MinMove(board, "X", 0)
		if score != -10 {
			t.Errorf("MinMove should choose winning move with score -10, got %d", score)
		}
	})
}