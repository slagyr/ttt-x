# TTT in Golang

Translated by ChatGPT (OpenAI) from Alex Root-Roatch's implementation: https://github.com/arootroatch/tic-tac-toe-clojure

## Basic Functionality

https://github.com/arootroatch/tic-tac-toe-clojure/tree/5a339a00bc23ca8dd9b4c14501a26dd4b4df8711

### Attempt 1

I asked openAI to translate the repo. The result was different to the original implemenation
1. The way it printed the board:
```
- - -
- - -
- - -

Player X, enter row and column (e.g., 1 2):  0 1

vs

1 2 3
4 5 6
7 8 9
Please enter your move (type 1-9 and hit enter):
3
```
2. The generated version was human vs human, whereas the original was human vs computer.  Big miss.
3. There was a test file but it was very minimal and did not pass.

It was almost like the AI tried to infer the intent of the original implementation without looking closely at the code.

### Attempt 2

Given the prompt:

    This should be human vs computer, try not to miss EVERY detail of this repository

ChatGPT added a computer player, but it's super simplistic, taking the first move available.  The output remains the same.
The test file still fails
the way the board was printed and how moves are inputted remained unchanged.


### Attemp 3

Given the prompt:
    
    This is not an unbeatable AI. Please make sure you are paying attention to every detail of this repository and copying it over as closely as you can in go.

ChatGPT made the game unbeatable but the game become automatically played out after the first turn.

```
- - -
- - -
- - -

Player X, enter row and column (e.g., 1 2): 0 0
X - -
- - -
- - -

Computer plays at row 0, col 1
X O X
O X O
X X X

Player X wins!
```

The game attempts to implement minimax but fails for two reasons:
1. the game plays out by itself after the first turn. It becomes AI vs. AI instead of player vs. AI
2. The algorithm seems to not work as it should because X won which should not be possible in a true minimax game.

Additionally, the response `"To create an unbeatable Tic-Tac-Toe AI that mimics the functionality of the Clojure repository using Go, we'll implement the Minimax algorithm."`
indicates that the AI pulled its own interpretation of minimax instead of using the implementation from the repository supplied.

The test file still fails
the way the board was printed and how moves are inputted remained unchanged.