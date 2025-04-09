package player

import (
	"testing"

	"ttt/board"
)

func TestMinimaxWinningMove(t *testing.T) {
	b := board.NewBoard()
	b.Cells = []string{"O", "O", "", "X", "X", "", "", "", ""}
	move := GetComputerMove(b)
	if move != 2 {
		t.Errorf("Expected move 2 to win with 'O', got %d", move)
	}
}

func TestMinimaxBlockingMove(t *testing.T) {
	b := board.NewBoard()
	b.Cells = []string{"X", "X", "", "", "O", "", "", "", ""}
	move := GetComputerMove(b)
	if move != 2 {
		t.Errorf("Expected move 2 to block 'X', got %d", move)
	}
}
