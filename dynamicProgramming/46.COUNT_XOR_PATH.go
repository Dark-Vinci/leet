package dynamicProgramming

func countPathsWithXorValue(grid [][]int, k int) int {
	var (
		dp   func(r, d, xor int, memo [][][16]int) int
		l    = len(grid)
		ll   = len(grid[0])
		MOD  = 1000000007
		memo = make([][][16]int, l)
	)

	for i := 0; i < l; i++ {
		memo[i] = make([][16]int, ll)
		for j := 0; j < ll; j++ {
			for k := 0; k < 16; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	dp = func(r, d, xor int, memo [][][16]int) int {
		if r == l || d == ll {
			return 0
		}

		if r == l-1 && d == ll-1 {
			if xor^grid[r][d] == k {
				return 1
			}

			return 0
		}

		if v := memo[r][d][xor]; v != -1 {
			return v
		}

		down := dp(r+1, d, xor^grid[r][d], memo)
		right := dp(r, d+1, xor^grid[r][d], memo)

		memo[r][d][xor] = (down + right) % MOD

		return memo[r][d][xor]
	}

	return dp(0, 0, 0, memo)
}
