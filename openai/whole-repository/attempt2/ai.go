package main

func GetBestMove(b *Board, player string) (int, int) {
	// Basic AI logic to pick the first available cell for simplicity
	for i := range b.cells {
		for j := range b.cells[i] {
			if b.cells[i][j] == "" {
				return i, j
			}
		}
	}
	// Shouldn't get here if the board is not full
	return -1, -1
}
