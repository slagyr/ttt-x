package board

// Board represents a Tic-Tac-Toe board with 9 cells.
type Board struct {
	Cells []string
}

// NewBoard creates a new empty 3x3 board.
func NewBoard() *Board {
	return &Board{Cells: make([]string, 9)}
}

// Mark places a mark ("X" or "O") at the given position.
func (b *Board) Mark(position int, mark string) {
	b.Cells[position] = mark
}

// IsPositionEmpty checks if a position is empty.
func (b *Board) IsPositionEmpty(position int) bool {
	return b.Cells[position] == ""
}

// IsFull checks if the board is full.
func (b *Board) IsFull() bool {
	for _, cell := range b.Cells {
		if cell == "" {
			return false
		}
	}
	return true
}

// GetWinner determines the winner ("X", "O", or "" if none).
func (b *Board) GetWinner() string {
	// Winning combinations: rows, columns, diagonals
	lines := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonals
	}
	for _, line := range lines {
		if b.Cells[line[0]] != "" &&
			b.Cells[line[0]] == b.Cells[line[1]] &&
			b.Cells[line[1]] == b.Cells[line[2]] {
			return b.Cells[line[0]]
		}
	}
	return ""
}

// AvailablePositions returns a list of empty position indices.
func (b *Board) AvailablePositions() []int {
	var positions []int
	for i, cell := range b.Cells {
		if cell == "" {
			positions = append(positions, i)
		}
	}
	return positions
}

// Copy creates a deep copy of the board.
func (b *Board) Copy() *Board {
	newCells := make([]string, 9)
	copy(newCells, b.Cells)
	return &Board{Cells: newCells}
}
