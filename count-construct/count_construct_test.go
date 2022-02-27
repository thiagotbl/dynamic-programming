package countconstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, CountConstruct("abcdef", []string{"ab", "abc", "def", "abcd"}))
	assert.Equal(2, CountConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	assert.Equal(
		0,
		CountConstruct(
			"skateboard",
			[]string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
		),
	)
	assert.Equal(
		4,
		CountConstruct(
			"enterapotentpot",
			[]string{"a", "p", "ent", "enter", "ot", "o", "t"},
		),
	)
	assert.Equal(
		0,
		CountConstruct(
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
