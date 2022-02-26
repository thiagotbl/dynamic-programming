package howsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestSum(t *testing.T) {
	assert := assert.New(t)

	assert.ElementsMatch(BestSum(7, []int{5, 3, 4, 7}), []int{7})
	assert.ElementsMatch(BestSum(8, []int{2, 3, 5}), []int{3, 5})
	assert.ElementsMatch(BestSum(8, []int{1, 4, 5}), []int{4, 4})
	assert.ElementsMatch(BestSum(100, []int{1, 2, 5, 25}), []int{25, 25, 25, 25})
}
