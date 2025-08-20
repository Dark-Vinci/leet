package dynamicProgramming

import "math"

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 0; i < len(dp); i++ {
		if dp[i] == math.MaxInt {
			continue
		}

		for _, op := range coins {
			if i+op < len(dp) {
				dp[i+op] = min(dp[i]+1, dp[i+op])
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}
