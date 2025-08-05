package dynamicProgramming

import "math"

func minFallingPathSum2(matrix [][]int) int {
	var (
		l      = len(matrix)
		ll     = len(matrix[0])
		memo   = make([][]int, l)
		result = math.MaxInt
		dp     func(i, j int, memo [][]int) int
	)

	for i := 0; i < l; i++ {
		memo[i] = make([]int, ll)
		for j := 0; j < ll; j++ {
			memo[i][j] = math.MaxInt
		}
	}

	dp = func(i, j int, memo [][]int) int {
		if i == l || j == ll {
			return math.MaxInt
		}

		if v := memo[i][j]; v != math.MaxInt {
			return v
		}

		if i == l-1 {
			memo[i][j] = matrix[i][j]
			return memo[i][j]
		}

		result := math.MaxInt

		for k := 0; k < ll; k++ {
			if k != j {
				result = min(result, dp(i+1, k, memo))
			}
		}

		memo[i][j] = matrix[i][j] + result

		return memo[i][j]
	}

	for j := 0; j < ll; j++ {
		result = min(result, dp(0, j, memo))
	}

	return result
}
