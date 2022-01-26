package fibonacci_memoized_struct

type fibonacciCalculator struct {
	values map[int]uint64
}

func (fc *fibonacciCalculator) Calculate(n int) uint64 {
	value, exists := fc.values[n]
	if exists {
		return value
	}

	fc.values[n] = fc.Calculate(n-1) + fc.Calculate(n-2)

	return fc.values[n]
}

func NewFibonacciCalculator() *fibonacciCalculator {
	return &fibonacciCalculator{
		values: map[int]uint64{
			0: 0,
			1: 1,
		},
	}
}
