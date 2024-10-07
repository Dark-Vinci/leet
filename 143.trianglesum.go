package main

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle)+1)

	for i := len(triangle) - 1; i >= 0; i-- {
		for j, val := range triangle[i] {
			dp[j] = val + min(dp[j+1], dp[j])
		}
	}

	return dp[0]
}
