package main

import (
	"fmt"
	"os"
	"strconv"
)

type fibonacciCalculator struct {
	values map[int]uint64
}

func (fc *fibonacciCalculator) calc(n int) uint64 {
	value, exists := fc.values[n]
	if exists {
		return value
	}

	fc.values[n] = fc.calc(n-1) + fc.calc(n-2)

	return fc.values[n]
}

func newFibonacciCalculator() *fibonacciCalculator {
	return &fibonacciCalculator{
		values: map[int]uint64{
			0: 0,
			1: 1,
		},
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("number not received")
	}

	input := os.Args[1]
	n, err := strconv.Atoi(input)
	if err != nil {
		panic("invalid input")
	}

	calculator := newFibonacciCalculator()

	fmt.Printf("f(%d) = %d\n", n, calculator.calc(n-1))
	fmt.Printf("f(%d) = %d\n", n, calculator.calc(n))
	fmt.Printf("f(%d) = %d\n", n, calculator.calc(n+1))
}
