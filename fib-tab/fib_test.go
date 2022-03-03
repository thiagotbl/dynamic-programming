package fibtab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert := assert.New(t)

	assert.EqualValues(0, Calculate(0))
	assert.EqualValues(1, Calculate(1))
	assert.EqualValues(1, Calculate(2))
	assert.EqualValues(2, Calculate(3))
	assert.EqualValues(3, Calculate(4))
	assert.EqualValues(5, Calculate(5))
	assert.EqualValues(55, Calculate(10))
	assert.EqualValues(610, Calculate(15))
	assert.EqualValues(12586269025, Calculate(50))
}
