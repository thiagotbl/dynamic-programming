package fibonacci_memoized

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciMemoized(t *testing.T) {
	assert := assert.New(t)

	assert.EqualValues(Calculate(0), 0)
	assert.EqualValues(Calculate(1), 1)
	assert.EqualValues(Calculate(2), 1)
	assert.EqualValues(Calculate(3), 2)
	assert.EqualValues(Calculate(4), 3)
	assert.EqualValues(Calculate(5), 5)
	assert.EqualValues(Calculate(10), 55)
	assert.EqualValues(Calculate(15), 610)
	assert.EqualValues(Calculate(20), 6765)
	assert.EqualValues(Calculate(40), 102334155)
}
