package canconstruct

import "strings"

// returns a boolean indicating whether or not the target
// can be constructed by concatenating elements fo the wordBank.
// elements of wordBank can be reused as many times as needed.
func CanConstruct(target string, wordBank []string) bool {
	memo := map[string]bool{}

	return canConstruct(target, wordBank, memo)
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if target == "" {
		return true
	}

	memoValue, exists := memo[target]
	if exists {
		return memoValue
	}

	for i := 0; i < len(wordBank); i++ {
		word := wordBank[i]

		if !strings.HasPrefix(target, word) {
			continue
		}

		newTarget := strings.Replace(target, word, "", 1)

		if canConstruct(newTarget, wordBank, memo) {
			memo[target] = true
			return true
		}
	}

	memo[target] = false
	return false
}
