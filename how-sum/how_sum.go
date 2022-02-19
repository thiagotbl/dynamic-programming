package howsum

// returns an array containing any combination of elements
// that add up to exactly the targetSum. If there is no combination
// that adds up to the targetSum, then return null.
//
// assumes all input members are nonnegative.
func HowSum(targetSum int, numbers []int) []int {
	memo := map[int][]int{}

	return howSum(targetSum, numbers, memo)
}

func howSum(targetSum int, numbers []int, memo map[int][]int) []int {
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

	for i := 0; i < len(numbers); i++ {
		newTarget := targetSum - numbers[i]
		newTargetResult := howSum(newTarget, numbers, memo)

		if newTargetResult != nil {
			targetSumResult := append(newTargetResult, numbers[i])

			memo[targetSum] = targetSumResult

			return targetSumResult
		}
	}

	memo[targetSum] = nil

	return nil
}
