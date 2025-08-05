package dynamicProgramming

import "math"

func minFallingPathSum(matrix [][]int) int {
	var (
		l      = len(matrix)
		ll     = len(matrix[0])
		memo   = make([][]int, l)
		dp     func(i, j int, memo [][]int) int
		result = math.MaxInt
	)

	for i := 0; i < l; i++ {
		memo[i] = make([]int, ll)
		for j := 0; j < ll; j++ {
			memo[i][j] = math.MaxInt
		}
	}

	dp = func(i, j int, memo [][]int) int {
		if i == l || j == ll || j < 0 {
			return math.MaxInt
		}

		if v := memo[i][j]; v != math.MaxInt {
			return v
		}

		if i == l-1 {
			memo[i][j] = matrix[i][j]
			return memo[i][j]
		}

		right, left := dp(i+1, j+1, memo), dp(i+1, j-1, memo)
		down := dp(i+1, j, memo)

		memo[i][j] = matrix[i][j] + min(right, down, left)

		return memo[i][j]
	}

	for j := 0; j < ll; j++ {
		result = min(result, dp(0, j, memo))
	}

	return result
}
