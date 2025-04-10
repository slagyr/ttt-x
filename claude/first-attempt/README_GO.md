# Tic-Tac-Toe in Go

This is a Go implementation of a command-line tic-tac-toe game with an unbeatable AI opponent.

## Features

- Command-line interface
- Play against an AI that uses the minimax algorithm
- Human plays as X, computer plays as O
- Color-coded board display
- Comprehensive test suite

## How to Run

Make sure you have Go installed on your system. Then:

1. Navigate to the project directory
2. Run the game:

```bash
go run main.go
```

## How to Play

- The board is numbered 1-9, with 1 at the top-left and 9 at the bottom-right
- When prompted, enter a number (1-9) to place your X in that position
- The computer will automatically make its move
- The game ends when someone wins or there's a tie

## Running Tests

The project includes a comprehensive test suite. To run all tests:

```bash
go test ./...
```

To run tests for a specific package:

```bash
go test ./board
go test ./player
```

To run a specific test with verbose output:

```bash
go test -v ./board -run TestScore
```

## Project Structure

- `main.go`: Main game loop and initialization
- `board/board.go`: Board representation and evaluation functions
- `player/player.go`: Player and AI logic, including the minimax algorithm
- `board/board_test.go`: Tests for board evaluation logic
- `player/player_test.go`: Tests for player and AI logic
- `board/print_test.go`: Tests for board printing functionality

## Implementation Notes

This Go implementation is a direct translation of the original Clojure tic-tac-toe application.
Key differences from the Clojure version include:

- Go's static typing vs. Clojure's dynamic typing
- Go's imperative programming style vs. Clojure's functional approach
- Go's struct-based data modeling vs. Clojure's vector-based approach
- Go's standard testing package vs. Clojure's speclj testing library