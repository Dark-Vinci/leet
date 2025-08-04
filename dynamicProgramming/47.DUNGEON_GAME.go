package dynamicProgramming

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	var (
		dp   func(r, d int, memo [][]int) int
		l    = len(dungeon)
		ll   = len(dungeon[0])
		memo = make([][]int, l)
	)

	for i := 0; i < l; i++ {
		memo[i] = make([]int, ll)
		for j := 0; j < ll; j++ {
			memo[i][j] = -1
		}
	}

	dp = func(r, d int, memo [][]int) int {
		if r == l || d == ll {
			return math.MaxInt
		}

		if r == l-1 && d == ll-1 {
			if dungeon[r][d] > 0 {
				return 1
			}

			return 1 - dungeon[r][d]
		}

		if v := memo[r][d]; v != -1 {
			return v
		}

		down, right := dp(r, d+1, memo), dp(r+1, d, memo)

		if min(down, right)-dungeon[r][d] <= 0 {
			memo[r][d] = 1
		} else {
			memo[r][d] = min(down, right) - dungeon[r][d]
		}

		return memo[r][d]
	}

	return dp(0, 0, memo)
}
