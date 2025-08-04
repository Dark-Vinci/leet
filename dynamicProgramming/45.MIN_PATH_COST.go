package dynamicProgramming

import "math"

func minPathCost(grid [][]int, moveCost [][]int) int {
	var (
		l    = len(grid)
		ll   = len(grid[0])
		memo = make([][]int, l+1)
		dp   func(i, j int, memo [][]int) int
	)

	for i := 0; i <= l; i++ {
		memo[i] = make([]int, ll+1)

		for j := 0; j <= ll; j++ {
			memo[i][j] = -1
		}
	}

	dp = func(i, j int, memo [][]int) int {
		if i == l-1 {
			return grid[i][j]
		}

		if v := memo[i+1][j+1]; v != -1 {
			return v
		}

		result := math.MaxInt

		for p := range ll {
			if i != -1 {
				result = min(result, dp(i+1, p, memo)+moveCost[grid[i][j]][p])
			} else {
				result = min(result, dp(i+1, p, memo))
			}
		}

		if i == -1 {
			memo[i+1][j+1] = result
		} else {
			memo[i+1][j+1] = result + grid[i][j]
		}

		return memo[i+1][j+1]
	}

	return dp(-1, -1, memo)
}
