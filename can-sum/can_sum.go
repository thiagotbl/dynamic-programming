package cansum

// returns a boolean indicating whether or not it is possible
// to generate the targetSum using members from the array.
//
// members can be used multiple times.
//
// assumes all input members are nonnegative.
func CanSum(targetSum int, numbers []int) bool {
	memo := map[int]bool{}

	return canSum(targetSum, numbers, memo)
}

func canSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	memoValue, exists := memo[targetSum]
	if exists {
		return memoValue
	}

	for i := 0; i < len(numbers); i++ {
		newTarget := targetSum - numbers[i]

		if canSum(newTarget, numbers, memo) {
			return true
		}

	}

	memo[targetSum] = false
	return false
}
