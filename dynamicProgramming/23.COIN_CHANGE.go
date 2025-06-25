package dynamicProgramming

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1

	for _, val := range coins {
		for i := val; i <= amount; i++ {
			dp[i] += dp[i-val]
		}
	}

	return dp[amount]
}
