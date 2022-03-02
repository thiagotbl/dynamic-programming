package allconstruct

import "strings"

// returns all of the ways that 'target' can be constructed
// concatenating elements of the 'wordBank'
func AllConstruct(target string, wordBank []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	var ways [][]string
	for i := 0; i < len(wordBank); i++ {
		word := wordBank[i]

		if !strings.HasPrefix(target, word) {
			continue
		}

		newTarget := strings.Replace(target, word, "", 1)
		newTargetWays := AllConstruct(newTarget, wordBank)

		for i := 0; i < len(newTargetWays); i++ {
			newTargetWays[i] = append([]string{word}, newTargetWays[i]...)
		}

		ways = append(ways, newTargetWays...)
	}

	return ways
}
