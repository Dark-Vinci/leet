package dynamicProgramming

import "cmp"

func hasValidPath(grid [][]byte) bool {
	var (
		l    = len(grid)
		ll   = len(grid[0])
		memo = make([][][]int, l)
		dp   func(i, j, slen int, memo [][][]int) bool
	)

	for i := 0; i < l; i++ {
		memo[i] = make([][]int, ll)
		for j := 0; j < ll; j++ {
			memo[i][j] = make([]int, l+ll)
			for k := 0; k < l+ll; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	dp = func(i, j, slen int, memo [][][]int) bool {
		if i == l || j == ll {
			return false
		}

		if grid[i][j] == '(' {
			slen++
		} else {
			slen--
		}

		if slen < 0 {
			return false
		}

		if v := memo[i][j][slen]; v != -1 {
			if v == 1 {
				return true
			} else {
				return false
			}
		}

		if i == l-1 && j == ll-1 {
			if slen == 0 {
				memo[i][j][slen] = 1
			} else {
				memo[i][j][slen] = 0
			}

			return slen == 0
		}

		down, right := dp(i+1, j, slen, memo), dp(i, j+1, slen, memo)

		result := cmp.Or(down, right)

		if result {
			memo[i][j][slen] = 1
		} else {
			memo[i][j][slen] = 0
		}

		return result
	}

	return dp(0, 0, 0, memo)
}
