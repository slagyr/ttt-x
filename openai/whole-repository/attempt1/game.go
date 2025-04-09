package main

import (
	"bytes"
)

type Board struct {
	cells [3][3]string
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) MakeMove(player string, row, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 || b.cells[row][col] != "" {
		return false
	}
	b.cells[row][col] = player
	return true
}

func (b *Board) IsFull() bool {
	for i := range b.cells {
		for j := range b.cells[i] {
			if b.cells[i][j] == "" {
				return false
			}
		}
	}
	return true
}

func (b *Board) CheckWinner() string {
	// Rows and Columns
	for i := 0; i < 3; i++ {
		if b.cells[i][0] != "" && b.cells[i][0] == b.cells[i][1] && b.cells[i][1] == b.cells[i][2] {
			return b.cells[i][0]
		}
		if b.cells[0][i] != "" && b.cells[0][i] == b.cells[1][i] && b.cells[1][i] == b.cells[2][i] {
			return b.cells[0][i]
		}
	}
	// Diagonals
	if b.cells[0][0] != "" && b.cells[0][0] == b.cells[1][1] && b.cells[1][1] == b.cells[2][2] {
		return b.cells[0][0]
	}
	if b.cells[0][2] != "" && b.cells[0][2] == b.cells[1][1] && b.cells[1][1] == b.cells[2][0] {
		return b.cells[0][2]
	}

	return ""
}

func (b *Board) String() string {
	var buf bytes.Buffer
	for i := range b.cells {
		for j, cell := range b.cells[i] {
			if cell == "" {
				buf.WriteString("-")
			} else {
				buf.WriteString(cell)
			}
			if j < 2 {
				buf.WriteString(" ")
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}
