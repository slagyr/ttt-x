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
	board := EmptyBoard()
	expected := Board{"x", "e", "e", "e", "e", "e", "e", "e", "e"}
	result := PlayerMove(board, 0, "x")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	board = PlayerMove(board, 0, "x")
	result = PlayerMove(board, 0, "o")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v (no overwrite), got %v", expected, result)
	}

	result = PlayerMove(board, 9, "x")
	if !reflect.DeepEqual(result, board) {
		t.Errorf("Expected %v (unchanged), got %v", board, result)
	}
}

func TestComputerMove(t *testing.T) {
	board := EmptyBoard()
	result := ComputerMove(board, "o")
	expected := Board{"o", "e", "e", "e", "e", "e", "e", "e", "e"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	board = Board{"x", "e", "e", "e", "e", "e", "e", "e", "e"}
	result = ComputerMove(board, "o")
	expected = Board{"x", "o", "e", "e", "e", "e", "e", "e", "e"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFullBoard(t *testing.T) {
	board := Board{"x", "o", "x", "e", "o", "x", "o", "x", "o"}
	if FullBoard(board) {
		t.Error("Expected board not full, but FullBoard returned true")
	}

	board = Board{"x", "o", "x", "o", "x", "o", "x", "o", "x"}
	if !FullBoard(board) {
		t.Error("Expected board full, but FullBoard returned false")
	}
}

func TestHasWinner(t *testing.T) {
	board := Board{"x", "x", "x", "e", "e", "e", "e", "e", "e"}
	if !HasWinner(board) {
		t.Error("Expected winner, but HasWinner returned false")
	}

	board = Board{"x", "e", "e", "x", "e", "e", "x", "e", "e"}
	if !HasWinner(board) {
		t.Error("Expected winner, but HasWinner returned false")
	}

	board = Board{"x", "e", "e", "e", "x", "e", "e", "e", "x"}
	if !HasWinner(board) {
		t.Error("Expected winner, but HasWinner returned false")
	}

	board = Board{"x", "o", "x", "e", "e", "e", "e", "e", "e"}
	if HasWinner(board) {
		t.Error("Expected no winner, but HasWinner returned true")
	}
}

func TestGameOver(t *testing.T) {
	board := Board{"x", "x", "x", "e", "e", "e", "e", "e", "e"}
	if !GameOver(board) {
		t.Error("Expected game over (win), but GameOver returned false")
	}

	board = Board{"x", "o", "x", "o", "x", "o", "o", "x", "o"}
	if !GameOver(board) {
		t.Error("Expected game over (draw), but GameOver returned false")
	}

	board = Board{"x", "o", "e", "e", "e", "e", "e", "e", "e"}
	if GameOver(board) {
		t.Error("Expected game not over, but GameOver returned true")
	}
}
