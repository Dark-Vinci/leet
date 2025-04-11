package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 1; i*i <= n; i++ {
		sq := i * i

		for j := sq; j <= n; j++ {
			dp[j] = min(dp[j-sq]+1, dp[j])
		}
	}

	return dp[n]
}

// time limit exceeded
func numSquares_(n int) int {
	ps := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		ps = append(ps, i*i)
	}

	result := math.MaxInt

	var dfs func(sum, count int)

	dfs = func(sum, count int) {
		if sum > n {
			return
		}

		if sum == n {
			result = min(result, count)
			return
		}

		for _, v := range ps {
			dfs(sum+v, count+1)
		}
	}

	dfs(0, 0)

	return result
}
