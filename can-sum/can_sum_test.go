package cansum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSum(t *testing.T) {
	assert := assert.New(t)

	assert.True(CanSum(7, []int{5, 3, 4, 7}))
	assert.True(CanSum(7, []int{2, 3}))
	assert.True(CanSum(8, []int{2, 3, 5}))
	assert.True(CanSum(9, []int{3, 34, 4, 12, 5, 2}))
	assert.True(CanSum(30, []int{3, 34, 4, 12, 5, 2}))

	assert.False(CanSum(2, []int{5, 8}))
	assert.False(CanSum(7, []int{2, 4}))
	assert.False(CanSum(300, []int{7, 14}))
}
