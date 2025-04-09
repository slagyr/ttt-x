package game

import (
	"testing"
)

func TestSwitchPlayer(t *testing.T) {
	if switchPlayer("X") != "O" {
		t.Error("Expected 'O' after 'X'")
	}
	if switchPlayer("O") != "X" {
		t.Error("Expected 'X' after 'O'")
	}
}
