package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		panic("number not received")
	}

	input := os.Args[1]
	n, err := strconv.Atoi(input)
	if err != nil {
		panic("invalid input")
	}

	result := fibonnaci(n)

	fmt.Printf("f(%d) = %d\n", n, result)
}

func fibonnaci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonnaci(n-1) + fibonnaci(n-2)
}
