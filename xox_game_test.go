package main

import "testing"

func TestMediumComputerMoveTakesWin(t *testing.T) {
	game := board{computer, computer, empty, human, human, empty, empty, empty, empty}

	if move := game.mediumComputerMove(); move != 2 {
		t.Fatalf("mediumComputerMove() = %d, want 2", move)
	}
}

func TestMediumComputerMoveBlocksHumanWin(t *testing.T) {
	game := board{human, human, empty, empty, computer, empty, empty, empty, empty}

	if move := game.mediumComputerMove(); move != 2 {
		t.Fatalf("mediumComputerMove() = %d, want 2", move)
	}
}

func TestComputerMoveChoosesEmptySquare(t *testing.T) {
	game := board{human, empty, empty, empty, computer, empty, empty, empty, empty}

	for _, level := range []difficulty{easy, medium, hard} {
		for range 10 {
			move := game.computerMove(level)
			if move < 0 || move >= len(game) {
				t.Fatalf("computerMove(%d) = %d, want a board position", level, move)
			}
			if game[move] != empty {
				t.Fatalf("computerMove(%d) = %d, want an empty square", level, move)
			}
		}
	}
}
