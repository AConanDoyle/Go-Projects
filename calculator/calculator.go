package main

import (
	"bufio"
	"fmt"
	"os"
)

var i, j int

func add() int {
	x := i + j
	return x
}

func minus() {
	fmt.Println("...")
}

func main() {
	fmt.Println("Welcome to a Calculator App.")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Scan(&i)
	fmt.Println("i: &i")
}
