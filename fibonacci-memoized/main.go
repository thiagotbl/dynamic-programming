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

	memo := map[int]uint64{0: 0, 1: 1}

	result := fibonnaciMemo(n, memo)

	fmt.Printf("f(%d) = %d\n", n, result)
}

func fibonnaciMemo(n int, memo map[int]uint64) uint64 {
	value, exists := memo[n]

	if exists {
		return value
	}

	memo[n] = fibonnaciMemo(n-1, memo) + fibonnaciMemo(n-2, memo)

	return memo[n]
}
