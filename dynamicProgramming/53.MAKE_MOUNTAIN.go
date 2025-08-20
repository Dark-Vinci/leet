package dynamicProgramming

import "math"

func makeMountain(nums []int) int {
	var (
		dp   func(prev, curr, state, inc, dec int, memo [][][3][2][2]int) int
		l    = len(nums)
		memo = make([][][3][2][2]int, l+1)
	)

	for i := 0; i <= l; i++ {
		memo[i] = make([][3][2][2]int, l)
		for j := 0; j < l; j++ {
			for k := 0; k < 3; k++ {
				for m := 0; m < 2; m++ {
					for n := 0; n < 2; n++ {
						memo[i][j][k][m][n] = -1
					}
				}
			}
		}
	}

	dp = func(prev, curr, state, inc, dec int, memo [][][3][2][2]int) int {
		if curr == l {
			if inc == 1 && dec == 1 {
				return 0
			}

			return math.MinInt
		}

		if v := memo[prev+1][curr][state][inc][dec]; v != -1 {
			return v
		}

		dont := dp(prev, curr+1, state, inc, dec, memo)
		do := math.MinInt

		if prev == -1 {
			do = 1 + dp(curr, curr+1, 1, 0, 0, memo)
		} else {
			if state == 1 {
				if nums[curr] > nums[prev] {
					do = 1 + dp(curr, curr+1, 1, 1, dec, memo)
				} else if nums[curr] < nums[prev] {
					do = 1 + dp(curr, curr+1, 2, inc, 1, memo)
				}
			} else if state == 2 {
				if nums[curr] < nums[prev] {
					do = 1 + dp(curr, curr+1, 2, inc, 1, memo)
				}
			}
		}

		memo[prev+1][curr][state][inc][dec] = max(dont, do)

		return memo[prev+1][curr][state][inc][dec]
	}

	return l - dp(-1, 0, 0, 0, 0, memo)
}
