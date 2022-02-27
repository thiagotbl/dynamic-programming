package countconstruct

import "strings"

// returns the number of ways the target can be constructed by
// concatenating elements of the wordbank array.
// elements of wordbank can be used as many times as needed.
func CountConstruct(target string, wordBank []string) int {
	memo := map[string]int{}

	return countConstruct(target, wordBank, memo)
}

func countConstruct(target string, wordBank []string, memo map[string]int) int {
	if target == "" {
		return 1
	}

	memoValue, exists := memo[target]
	if exists {
		return memoValue
	}

	ways := 0
	for i := 0; i < len(wordBank); i++ {
		word := wordBank[i]

		if !strings.HasPrefix(target, word) {
			continue
		}

		newTarget := strings.Replace(target, word, "", 1)

		ways += countConstruct(newTarget, wordBank, memo)
	}

	memo[target] = ways
	return ways
}
