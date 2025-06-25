package dynamicProgramming

import "math"

// what i ended up with
func minimumCoins(prices []int) int {
	var dp func(i int, memo map[int]int) int

	dp = func(i int, memo map[int]int) int {
		if i >= len(prices) {
			return 0
		}

		if v, ok := memo[i]; ok {
			return v
		}

		mn, mj := math.MaxInt, min(2*(i+1), len(prices))

		for j := i + 1; j <= mj; j++ {
			curr := dp(j, memo)

			mn = min(mn, curr)
		}

		memo[i] = mn + prices[i]

		return memo[i]
	}

	return dp(0, map[int]int{})
}

// others; more concise solution
func minimumCoins_(prices []int) int {
	dp := make([]int, len(prices)+1)
	for i := len(prices) - 1; i >= 0; i-- {
		dp[i] = math.MaxInt32
		for j := i + 1; j <= i+i+2 && j <= len(prices); j++ {
			dp[i] = min(dp[i], prices[i]+dp[j])
		}
	}
	return dp[0]
}

// others; more concise solution another one
func minimumCoins__(prices []int) int {
	n := len(prices)
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt
	}

	dp[n] = 0

	for i := n - 1; i > -1; i-- {
		for j := i + 1; j < min(i+i+3, n+1); j++ {
			dp[i] = min(dp[i], dp[j]+prices[i])
		}
	}

	return dp[0]
}
