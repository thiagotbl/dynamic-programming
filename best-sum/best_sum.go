package howsum

// returns an array containing the shortest combination
// of numbers that add up to exactly the targetSum.
// if there is a tie for the shortest combination, return any of the sortest.
func BestSum(targetSum int, numbers []int) []int {
	memo := map[int][]int{}

	return bestSum(targetSum, numbers, memo)
}

func bestSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	memoValue, exists := memo[targetSum]
	if exists {
		return memoValue
	}

	var shortestResult []int
	for i := 0; i < len(numbers); i++ {
		newTarget := targetSum - numbers[i]
		newTargetResult := bestSum(newTarget, numbers, memo)

		if newTargetResult == nil {
			continue
		}

		newTargetResult = append(newTargetResult, numbers[i])

		isNewShorter := shortestResult == nil ||
			len(newTargetResult) < len(shortestResult)

		if isNewShorter {
			shortestResult = newTargetResult
		}
	}

	memo[targetSum] = shortestResult
	return shortestResult
}
