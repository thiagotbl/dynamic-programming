package grid_traveler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert := assert.New(t)

	assert.EqualValues(GridTraveler(1, 1), 1)
	assert.EqualValues(GridTraveler(2, 2), 2)
	assert.EqualValues(GridTraveler(2, 3), 3)
	assert.EqualValues(GridTraveler(3, 2), 3)
	assert.EqualValues(GridTraveler(3, 3), 6)
	assert.EqualValues(GridTraveler(18, 18), 2333606220)
}
