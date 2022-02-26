package howsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	assert := assert.New(t)

	assert.True(CanConstruct("abcdef", []string{"ab", "abc", "def", "abcd"}))
	assert.True(CanConstruct("", []string{"cat", "dog", "mouse"}))

	assert.False(CanConstruct(
		"saketeboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}),
	)
	assert.False(
		CanConstruct(
			"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
			[]string{
				"e",
				"ee",
				"eee",
				"eeee",
				"eeeeee",
				"eeeeee",
			},
		),
	)
}
