package main

import (
	"fmt"
)

var i, j int

func add() int {
	x := i + j
	fmt.Println("Result:")
	return x
}

func minus() {
	fmt.Println("Result:")
}

func main() {
	fmt.Scan(&i)
	fmt.Println("Result:")
}
