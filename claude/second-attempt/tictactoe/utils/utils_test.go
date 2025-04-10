package utils

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"tictactoe/board"
)

// Helper function to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	f()
	
	w.Close()
	os.Stdout = old
	
	var buf bytes.Buffer
	io.Copy(&buf, r)
	
	return buf.String()
}

func TestDisplayInvalidMoveError(t *testing.T) {
	output := captureOutput(func() {
		DisplayInvalidMoveError()
	})
	
	expected := "Please choose an empty cell\n"
	if output != expected {
		t.Errorf("DisplayInvalidMoveError() output = %q, expected %q", output, expected)
	}
}

func TestPromptUserForMove(t *testing.T) {
	output := captureOutput(func() {
		PromptUserForMove()
	})
	
	expected := "Please enter your move (type 1-9 and hit enter):\n"
	if output != expected {
		t.Errorf("PromptUserForMove() output = %q, expected %q", output, expected)
	}
}

func TestDisplayBotMoveMessage(t *testing.T) {
	output := captureOutput(func() {
		DisplayBotMoveMessage(5)
	})
	
	expected := "Your opponent plays 5!\n"
	if output != expected {
		t.Errorf("DisplayBotMoveMessage() output = %q, expected %q", output, expected)
	}
}

func TestPrintBoard(t *testing.T) {
	// Test with a mixed board
	board := []interface{}{1, board.X, 3, 4, board.O, 6, board.O, 8, 9}
	
	output := captureOutput(func() {
		PrintBoard(board)
	})
	
	// Check if output contains purple color codes for numbers
	purplePrefix := "\u001b[35m"
	resetSuffix := "\u001b[0m"
	
	// Check that all numbers have color formatting
	for _, num := range []int{1, 3, 4, 6, 8, 9} {
		coloredNum := purplePrefix + string(rune('0'+num)) + resetSuffix
		if !strings.Contains(output, coloredNum) {
			t.Errorf("PrintBoard() output doesn't contain colored number %d: %q", num, output)
		}
	}
	
	// Check that X and O are displayed correctly (without color)
	if !strings.Contains(output, "X") || !strings.Contains(output, "O") {
		t.Errorf("PrintBoard() output doesn't contain player marks correctly: %q", output)
	}
	
	// Check that we have 3 lines for the board rows
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 3 {
		t.Errorf("PrintBoard() expected 3 lines of output, got %d: %q", len(lines), output)
	}
}