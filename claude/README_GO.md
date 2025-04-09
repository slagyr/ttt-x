# Tic-Tac-Toe in Go

This is a Go implementation of a command-line tic-tac-toe game with an unbeatable AI opponent.

## Features

- Command-line interface
- Play against an AI that uses the minimax algorithm
- Human plays as X, computer plays as O
- Color-coded board display

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

## Project Structure

- `main.go`: Main game loop and initialization
- `board/board.go`: Board representation and evaluation functions
- `player/player.go`: Player and AI logic, including the minimax algorithm

## Implementation Notes

This Go implementation is a direct translation of the original Clojure tic-tac-toe application.
Key differences from the Clojure version include:

- Go's static typing vs. Clojure's dynamic typing
- Go's imperative programming style vs. Clojure's functional approach
- Go's struct-based data modeling vs. Clojure's vector-based approach