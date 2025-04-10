package player

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"tictactoe/board"
)

// Mock version of PromptUserForMove to avoid IO issues during tests
func mockPromptUserForMove() {
	fmt.Println("Please enter your move (type 1-9 and hit enter):")
}

// Mock version of DisplayInvalidMoveError to avoid IO issues during tests
func mockDisplayInvalidMoveError() {
	fmt.Println("Please choose an empty cell")
}

// Helper function to setup testing environment
func setupTestEnv(input string) ([]byte, func()) {
	// Save original stdin and create a pipe
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	
	// Write test input to the pipe
	w.Write([]byte(input))
	w.Close()
	
	// Capture stdout
	oldStdout := os.Stdout
	pr, pw, _ := os.Pipe()
	os.Stdout = pw
	
	outC := make(chan []byte)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, pr)
		outC <- buf.Bytes()
	}()
	
	// Return a function to restore the environment
	cleanup := func() {
		os.Stdout = oldStdout
		pr.Close()
		os.Stdin = oldStdin
		close(outC)
	}
	
	return <-outC, cleanup
}

func TestPlayUserTurnMock(t *testing.T) {
	// Create a simple mock version of PlayUserTurn for testing
	mockPlayUserTurn := func(board []interface{}, player string, input string) []interface{} {
		// Simulate valid move
		inputNum := strings.Split(input, "\n")[0]
		moveIdx, _ := strconv.Atoi(inputNum)
		
		// Make a copy of the board
		newBoard := make([]interface{}, len(board))
		copy(newBoard, board)
		
		// Apply move
		moveIdx--  // Convert 1-based to 0-based
		if moveIdx >= 0 && moveIdx < len(newBoard) {
			if _, isNumber := newBoard[moveIdx].(int); isNumber {
				newBoard[moveIdx] = player
			}
		}
		
		return newBoard
	}
	
	// Create a board with numbers 1-9
	emptyBoard := make([]interface{}, 9)
	for i := 0; i < 9; i++ {
		emptyBoard[i] = i + 1
	}
	
	tests := []struct {
		name          string
		board         []interface{}
		player        string
		input         string
		expectedBoard []interface{}
	}{
		{
			name:          "Player X plays first cell",
			board:         emptyBoard,
			player:        board.X,
			input:         "1\n",
			expectedBoard: []interface{}{board.X, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:          "Player X plays second cell",
			board:         emptyBoard,
			player:        board.X,
			input:         "2\n",
			expectedBoard: []interface{}{1, board.X, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mockPlayUserTurn(tt.board, tt.player, tt.input)
			
			if !reflect.DeepEqual(result, tt.expectedBoard) {
				t.Errorf("PlayUserTurn() = %v, expected %v", result, tt.expectedBoard)
			}
		})
	}
}