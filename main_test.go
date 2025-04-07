package main

import (
	"reflect"
	"testing"
)

func TestEmptyBoard(t *testing.T) {
	board := EmptyBoard()
	if len(board) != 9 {
		t.Errorf("Expected board length 9, got %d", len(board))
	}
	expected := Board{"e", "e", "e", "e", "e", "e", "e", "e", "e"}
	if !reflect.DeepEqual(board, expected) {
		t.Errorf("Expected %v, got %v", expected, board)
	}
}

func TestPlayerMove(t *testing.T) {
	// Place mark on empty space
	board := EmptyBoard()
	expected := Board{"x", "e", "e", "e", "e", "e", "e", "e", "e"}
	result := PlayerMove(board, 0, "x")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Don't overwrite existing mark
	board = PlayerMove(board, 0, "x")
	result = PlayerMove(board, 0, "o")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v (no overwrite), got %v", expected, result)
	}

	// Invalid position
	result = PlayerMove(board, 9, "x")
	if !reflect.DeepEqual(result, board) {
		t.Errorf("Expected %v (unchanged), got %v", board, result)
	}
}

func TestFullBoard(t *testing.T) {
	// Not full
	board := Board{"x", "o", "x", "e", "o", "x", "o", "x", "o"}
	if FullBoard(board) {
		t.Error("Expected board not full, but FullBoard returned true")
	}

	// Full
	board = Board{"x", "o", "x", "o", "x", "o", "x", "o", "x"}
	if !FullBoard(board) {
		t.Error("Expected board full, but FullBoard returned false")
	}
}

func TestHasWinner(t *testing.T) {
	// Horizontal win
	board := Board{"x", "x", "x", "e", "e", "e", "e", "e", "e"}
	if !HasWinner(board) {
		t.Error("Expected winner, but HasWinner returned false")
	}

	// Vertical win
	board = Board{"x", "e", "e", "x", "e", "e", "x", "e", "e"}
	if !HasWinner(board) {
		t.Error("Expected winner, but HasWinner returned false")
	}

	// Diagonal win
	board = Board{"x", "e", "e", "e", "x", "e", "e", "e", "x"}
	if !HasWinner(board) {
		t.Error("Expected winner, but HasWinner returned false")
	}

	// No win
	board = Board{"x", "o", "x", "e", "e", "e", "e", "e", "e"}
	if HasWinner(board) {
		t.Error("Expected no winner, but HasWinner returned true")
	}
}

func TestGameOver(t *testing.T) {
	// Win
	board := Board{"x", "x", "x", "e", "e", "e", "e", "e", "e"}
	if !GameOver(board) {
		t.Error("Expected game over (win), but GameOver returned false")
	}

	// Draw
	board = Board{"x", "o", "x", "o", "x", "o", "o", "x", "o"}
	if !GameOver(board) {
		t.Error("Expected game over (draw), but GameOver returned false")
	}

	// Not over
	board = Board{"x", "o", "e", "e", "e", "e", "e", "e", "e"}
	if GameOver(board) {
		t.Error("Expected game not over, but GameOver returned true")
	}
}
