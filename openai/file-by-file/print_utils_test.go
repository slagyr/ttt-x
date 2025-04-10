// test_utils_test.go - Tests for the print utilities in the tic_tac_toe package.
package tic_tac_toe

import (
	"bytes"
	"os"
	"testing"
)

// helper function to capture output

func captureOutput(f func()) string {
	// Save the original stdout
	old := os.Stdout

	// Create a pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	f()

	// Close writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.String()
}

// TestPrintBoard tests the printBoard function.
func TestPrintBoard(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{" ", "X", "O"},
		{"O", " ", " "},
	}

	// Capture the output of printBoard
	output := captureOutput(func() {
		printBoard(board)
	})

	expectedOutput := "[X O X]\n[  X O]\n[O    ]\n"

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\n but got:\n%s", expectedOutput, output)
	}
}

// TestPrintInstructions tests the printInstructions function.
func TestPrintInstructions(t *testing.T) {
	// Assuming that the instructions should print a specific string
	expected := "Welcome to Tic-Tac-Toe!\nHere are the instructions...\n"

	// Capture the output of printInstructions
	output := captureOutput(printInstructions)

	if output != expected {
		t.Errorf("Expected %s but got %s", expected, output)
	}
}
