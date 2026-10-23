package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	empty    byte = ' '
	human    byte = 'X'
	computer byte = 'O'
)

type difficulty int

const (
	easy difficulty = iota + 1
	medium
	hard
)

var winningLines = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

type board [9]byte

func (b board) winner() byte {
	for _, line := range winningLines {
		if b[line[0]] != empty && b[line[0]] == b[line[1]] && b[line[1]] == b[line[2]] {
			return b[line[0]]
		}
	}
	return empty
}

func (b board) full() bool {
	for _, cell := range b {
		if cell == empty {
			return false
		}
	}
	return true
}

func (b board) print() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if col > 0 {
				fmt.Print(" | ")
			}
			cell := b[row*3+col]
			if cell == empty {
				fmt.Print(row*3 + col + 1)
			} else {
				fmt.Printf("%c", cell)
			}
		}
		fmt.Println()
		if row < 2 {
			fmt.Println("--+---+--")
		}
	}
}

func minimax(b board, isComputerTurn bool, depth int) int {
	switch b.winner() {
	case computer:
		return 10 - depth
	case human:
		return depth - 10
	}
	if b.full() {
		return 0
	}

	bestScore := 100
	if isComputerTurn {
		bestScore = -100
	}
	for position, cell := range b {
		if cell != empty {
			continue
		}
		if isComputerTurn {
			b[position] = computer
			score := minimax(b, false, depth+1)
			if score > bestScore {
				bestScore = score
			}
		} else {
			b[position] = human
			score := minimax(b, true, depth+1)
			if score < bestScore {
				bestScore = score
			}
		}
		b[position] = empty
	}
	return bestScore
}

func (b board) bestComputerMove() int {
	bestPosition := -1
	bestScore := -100
	for position, cell := range b {
		if cell != empty {
			continue
		}
		b[position] = computer
		score := minimax(b, false, 1)
		b[position] = empty
		if score > bestScore {
			bestScore = score
			bestPosition = position
		}
	}
	return bestPosition
}

func (b board) randomComputerMove() int {
	var moves []int
	for position, cell := range b {
		if cell == empty {
			moves = append(moves, position)
		}
	}
	if len(moves) == 0 {
		return -1
	}
	return moves[rand.Intn(len(moves))]
}

func (b board) winningMove(mark byte) int {
	for position, cell := range b {
		if cell != empty {
			continue
		}
		next := b
		next[position] = mark
		if next.winner() == mark {
			return position
		}
	}
	return -1
}

func (b board) mediumComputerMove() int {
	if position := b.winningMove(computer); position >= 0 {
		return position
	}
	if position := b.winningMove(human); position >= 0 {
		return position
	}

	if b[4] == empty {
		return 4
	}

	var corners []int
	for _, position := range []int{0, 2, 6, 8} {
		if b[position] == empty {
			corners = append(corners, position)
		}
	}
	if len(corners) > 0 {
		return corners[rand.Intn(len(corners))]
	}
	return b.randomComputerMove()
}

func (b board) computerMove(level difficulty) int {
	switch level {
	case easy:
		return b.randomComputerMove()
	case medium:
		return b.mediumComputerMove()
	default:
		return b.bestComputerMove()
	}
}

func main() {
	var game board
	for position := range game {
		game[position] = empty
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Tic-Tac-Toe: you are X, the computer is O.")
	fmt.Println("Choose difficulty: 1 = easy, 2 = medium, 3 = hard.")
	var level difficulty
	for level == 0 {
		fmt.Print("Difficulty (1-3): ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Input error: %v\n", err)
			}
			fmt.Println("Game ended.")
			return
		}
		switch strings.TrimSpace(scanner.Text()) {
		case "1":
			level = easy
		case "2":
			level = medium
		case "3":
			level = hard
		default:
			fmt.Println("Enter 1 for easy, 2 for medium, or 3 for hard.")
		}
	}
	fmt.Println("Choose a position from 1 to 9:")

	for {
		game.print()
		fmt.Print("Your move (1-9): ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Input error: %v\n", err)
			}
			fmt.Println("Game ended.")
			return
		}

		position, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || position < 1 || position > 9 {
			fmt.Println("Enter a number from 1 to 9.")
			continue
		}
		position--
		if game[position] != empty {
			fmt.Println("That position is already taken.")
			continue
		}

		game[position] = human
		if game.winner() == human {
			game.print()
			fmt.Println("You win!")
			return
		}
		if game.full() {
			game.print()
			fmt.Println("It's a draw!")
			return
		}

		computerPosition := game.computerMove(level)
		game[computerPosition] = computer
		fmt.Printf("Computer chooses position %d.\n", computerPosition+1)
		if game.winner() == computer {
			game.print()
			fmt.Println("The computer wins.")
			return
		}
		if game.full() {
			game.print()
			fmt.Println("It's a draw!")
			return
		}
	}
}
