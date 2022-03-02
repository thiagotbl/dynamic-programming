package allconstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllConstruct(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[][]string{{}},
		AllConstruct("", []string{"a", "b"}),
	)

	assert.Equal(
		[][]string{{}},
		AllConstruct("", []string{"cat", "dog", "mouse"}),
	)

	assert.Empty(
		AllConstruct("purple", []string{"a", "b", "c"}),
	)

	assert.Empty(
		AllConstruct("hello", []string{"cat", "dog", "mouse"}),
	)

	assert.Equal(
		[][]string{
			{"purp", "le"},
			{"p", "ur", "p", "le"},
		},
		AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}),
	)

	assert.Equal(
		[][]string{
			{"ab", "cd", "ef"},
			{"ab", "c", "def"},
			{"abc", "def"},
			{"abcd", "ef"},
		},
		AllConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}),
	)
}
