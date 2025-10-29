package stacks_arrays

import "cmp"

// CORRECT
func isSubsequence_(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	var (
		dp   func(curr, ind int, memo [][]int) bool
		memo = make([][]int, len(s))
	)

	for i := 0; i < len(s); i++ {
		memo[i] = make([]int, len(t))

		for j := 0; j < len(t); j++ {
			memo[i][j] = -1
		}
	}

	dp = func(curr, ind int, memo [][]int) bool {
		if curr >= len(t) && ind != len(s) {
			return false
		}

		if ind == len(s) {
			return true
		}

		if memo[ind][curr] != -1 {
			if memo[ind][curr] == 0 {
				return false
			} else {
				return true
			}
		}

		dont := dp(curr+1, ind, memo)

		if s[ind] == t[curr] {
			dont = cmp.Or(dont, dp(curr+1, ind+1, memo))
		}

		if dont == false {
			memo[ind][curr] = 0
		} else {
			memo[ind][curr] = 1
		}

		return dont
	}

	return dp(0, 0, memo)
}

// CORRECT
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	curr := 0

	for i := 0; i < len(t); i++ {
		if curr < len(s) && t[i] == s[curr] {
			curr++
		}
	}

	return len(s) == curr
}
