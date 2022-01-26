package fibonacci_memoized

func Calculate(n int) uint64 {
	memo := map[int]uint64{0: 0, 1: 1}

	return fibonacciMemo(n, memo)
}

func fibonacciMemo(n int, memo map[int]uint64) uint64 {
	value, exists := memo[n]

	if exists {
		return value
	}

	memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)

	return memo[n]
}
