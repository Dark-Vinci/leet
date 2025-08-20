package stacks_arrays

func maxProfit(prices []int) int {
	minimum, result := prices[0], 0

	for i := 0; i < len(prices); i++ {
		result = max(result, prices[i]-minimum)
		minimum = min(minimum, prices[i])
	}

	return result
}
