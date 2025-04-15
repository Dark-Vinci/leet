package dynamicProgramming

import (
	"math"
	"slices"
)

func maxPoints(points [][]int) int64 {
	dp := make([][]int64, len(points))

	l, ll := len(points), len(points[0])

	for i := 0; i < l; i++ {
		dp[i] = make([]int64, ll)
	}

	for i := 0; i < ll; i++ {
		dp[0][i] = int64(points[0][i])
	}

	for i := 1; i < l; i++ {
		lr, rl := make([]int64, ll), make([]int64, ll)

		lr[0] = dp[i-1][0]
		rl[ll-1] = dp[i-1][ll-1]

		for j := 1; j < ll; j++ {
			lr[j] = max(lr[j-1]-1, dp[i-1][j])
			rl[ll-j-1] = max(rl[ll-j]-1, dp[i-1][ll-j-1])
		}

		for j := 0; j < ll; j++ {
			dp[i][j] = int64(points[i][j]) + max(lr[j], rl[j])
		}
	}

	return slices.Max(dp[len(dp)-1])
}

func maxPoints_(points [][]int) int64 {
	var helper func(i int, lIndex int, sum int64) int64

	helper = func(i int, lIndex int, sum int64) int64 {
		if i >= len(points) {
			return sum
		}

		mx, col := math.MinInt, -1

		for j, v := range points[i] {
			if i == 0 {
				if v >= mx {
					mx = v
					col = j
				}
				continue
			}

			cDiff := abs(j - lIndex)
			value := v - cDiff

			if value >= mx {
				mx = value
				col = j
			}
		}

		return helper(i+1, col, sum+int64(mx))
	}

	return helper(0, 0, 0)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
