package howsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHowSum(t *testing.T) {
	assert := assert.New(t)

	assert.ElementsMatch(HowSum(7, []int{2, 3}), []int{3, 2, 2})
	assert.ElementsMatch(HowSum(7, []int{5, 3, 4, 7}), []int{4, 3})

	assert.Nil(HowSum(7, []int{2, 4}))

	assert.ElementsMatch(HowSum(8, []int{2, 3, 5}), []int{2, 2, 2, 2})

	assert.Nil(HowSum(300, []int{7, 14}))
}
