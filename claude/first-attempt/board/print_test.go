package board

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Test PrintBoard function
func TestPrintBoard(t *testing.T) {
	// Redirect stdout to capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	// Create test board
	board := Board{
		"1", "X", "3", 
		"4", "O", "6", 
		"O", "8", "9",
	}
	
	// Call the function we want to test
	PrintBoard(board)
	
	// Restore stdout
	w.Close()
	os.Stdout = old
	
	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()
	
	// Verify output contains the board elements
	// (We can't test the exact formatting due to color codes, but we can check content)
	if !strings.Contains(output, "1") || !strings.Contains(output, "X") || 
	   !strings.Contains(output, "3") || !strings.Contains(output, "4") ||
	   !strings.Contains(output, "O") || !strings.Contains(output, "6") ||
	   !strings.Contains(output, "O") || !strings.Contains(output, "8") ||
	   !strings.Contains(output, "9") {
		t.Errorf("PrintBoard output does not contain expected board elements")
	}
	
	// Verify the newlines for rows
	lines := strings.Split(output, "\n")
	if len(lines) < 3 {
		t.Errorf("PrintBoard output should have at least 3 lines, got %d", len(lines))
	}
}