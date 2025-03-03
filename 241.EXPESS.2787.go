package main

func numberOfWays(n int, x int) int {
	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		p := 1

		for j := 0; j < x; j++ {
			p *= i
		}

		for k := n; k >= p; k-- {
			dp[k] += dp[k-p]
			dp[k] %= 1_000_000_007
		}
	}

	return dp[n]
}
