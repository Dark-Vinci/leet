package dynamicProgramming

func maxValueOfCoins(piles [][]int, k int) int {
	var (
		n    = len(piles)
		dp   func(i, kk int, memo [][]int) int
		memo = make([][]int, n)
	)

	for i := 0; i < n; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			memo[i][j] = -1
		}
	}

	dp = func(i, kk int, memo [][]int) int {
		if i == n {
			return 0
		}

		if v := memo[i][kk]; v != -1 {
			return v
		}

		dont := dp(i+1, kk, memo)

		sum, result := 0, 0

		for j := 0; j < min(len(piles[i]), kk); j++ {
			sum += piles[i][j]
			result = max(result, sum+dp(i+1, kk-j-1, memo))
		}

		memo[i][kk] = max(result, dont)

		return memo[i][kk]
	}

	return dp(0, k, memo)
}
