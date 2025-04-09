package main

import "testing"

func TestWinConditions(t *testing.T) {
	tests := []struct {
		setup  func(*Board)
		winner string
	}{
		{
			setup: func(b *Board) {
				b.MakeMove("X", 0, 0)
				b.MakeMove("X", 0, 1)
				b.MakeMove("X", 0, 2)
			},
			winner: "X",
		},
		{
			setup: func(b *Board) {
				b.MakeMove("O", 0, 0)
				b.MakeMove("O", 1, 0)
				b.MakeMove("O", 2, 0)
			},
			winner: "O",
		},
		{
			setup: func(b *Board) {
				b.MakeMove("X", 0, 2)
				b.MakeMove("X", 1, 1)
				b.MakeMove("X", 2, 0)
			},
			winner: "",
		},
	}

	for _, test := range tests {
		board := NewBoard()
		test.setup(board)
		if got := board.CheckWinner(); got != test.winner {
			t.Errorf("Expected winner %s, but got %s", test.winner, got)
		}
	}
}
