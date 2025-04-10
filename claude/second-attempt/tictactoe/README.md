# Tic-Tac-Toe

A command-line Tic-Tac-Toe game implemented in Go.

## Features

- Play against an AI opponent that uses the minimax algorithm
- Colorized output to make moves easier to see
- Simple command-line interface

## Project Structure

```
tictactoe/
├── board/          # Board evaluation and game state logic
├── bot/            # AI opponent using minimax algorithm  
├── cmd/            # Command-line application
│   └── tictactoe/  # Main application entry point
├── player/         # Human player move logic
└── utils/          # Formatting and display utilities
```

## Getting Started

To build and run the game:

```bash
cd tictactoe
go build ./cmd/tictactoe
./tictactoe
```

Or run directly with:

```bash
go run ./cmd/tictactoe
```

## Running Tests

```bash
go test ./...
```

## Game Rules

- Player X (you) plays first, followed by Player O (computer)
- Choose a cell by typing its number (1-9) and pressing Enter
- The first player to get three marks in a row (horizontally, vertically, or diagonally) wins
- If the board fills up without a winner, the game ends in a tie