package fibonacci

func Calculate(n int) uint64 {
	return fibonacci(uint64(n))
}

func fibonacci(n uint64) uint64 {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
