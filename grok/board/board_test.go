package board

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	if len(b.Cells) != 9 {
		t.Errorf("Expected 9 cells, got %d", len(b.Cells))
	}
	for i, cell := range b.Cells {
		if cell != "" {
			t.Errorf("Expected empty cell at %d, got %s", i, cell)
		}
	}
}

func TestMark(t *testing.T) {
	b := NewBoard()
	b.Mark(4, "X")
	if b.Cells[4] != "X" {
		t.Errorf("Expected 'X' at position 4, got %s", b.Cells[4])
	}
}

func TestIsPositionEmpty(t *testing.T) {
	b := NewBoard()
	if !b.IsPositionEmpty(0) {
		t.Error("Expected position 0 to be empty")
	}
	b.Mark(0, "O")
	if b.IsPositionEmpty(0) {
		t.Error("Expected position 0 to be occupied")
	}
}

func TestIsFull(t *testing.T) {
	b := NewBoard()
	if b.IsFull() {
		t.Error("Empty board should not be full")
	}
	for i := 0; i < 9; i++ {
		b.Mark(i, "X")
	}
	if !b.IsFull() {
		t.Error("Filled board should be full")
	}
}

func TestGetWinner(t *testing.T) {
	b := NewBoard()
	b.Cells = []string{"X", "X", "X", "", "", "", "", "", ""}
	if b.GetWinner() != "X" {
		t.Errorf("Expected winner 'X', got %s", b.GetWinner())
	}
	b = NewBoard()
	b.Cells = []string{"O", "", "", "O", "", "", "O", "", ""}
	if b.GetWinner() != "O" {
		t.Errorf("Expected winner 'O', got %s", b.GetWinner())
	}
	b = NewBoard()
	if b.GetWinner() != "" {
		t.Errorf("Expected no winner, got %s", b.GetWinner())
	}
}
