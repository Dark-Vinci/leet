package dynamicProgramming

import "math"

func maximumEnergy(energy []int, k int) int {
	var (
		memo   = make(map[int]int)
		dp     func(i int, memo map[int]int) int
		result = math.MinInt
	)

	dp = func(i int, memo map[int]int) int {
		if i >= len(energy) {
			return 0
		}

		if v, ok := memo[i]; ok {
			return v
		}

		memo[i] = energy[i] + dp(i+k, memo)

		return memo[i]
	}

	for i := 0; i < len(energy); i++ {
		result = max(result, dp(i, memo))
	}

	return result
}
