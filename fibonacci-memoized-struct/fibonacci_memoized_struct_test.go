package fibonacci_memoized_struct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciMemoizedStruct(t *testing.T) {
	assert := assert.New(t)

	calculator := NewFibonacciCalculator()

	assert.EqualValues(calculator.Calculate(0), 0)
	assert.EqualValues(calculator.Calculate(1), 1)
	assert.EqualValues(calculator.Calculate(2), 1)
	assert.EqualValues(calculator.Calculate(3), 2)
	assert.EqualValues(calculator.Calculate(4), 3)
	assert.EqualValues(calculator.Calculate(5), 5)
	assert.EqualValues(calculator.Calculate(10), 55)
	assert.EqualValues(calculator.Calculate(15), 610)
	assert.EqualValues(calculator.Calculate(20), 6765)
	assert.EqualValues(calculator.Calculate(40), 102334155)
}
