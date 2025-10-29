package arrays

import "math"

func longestMountain(nums []int) int {
	var (
		l      = len(nums)
		dp     = [2][]int{}
		result = 0
	)

	dp[0] = make([]int, l)
	dp[1] = make([]int, l)

	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			dp[0][i] = dp[0][i-1] + 1
		}
	}

	for i := l - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			dp[1][i] = dp[1][i+1] + 1
		}
	}

	for i := 0; i < l; i++ {
		if dp[0][i] > 0 && dp[1][i] > 0 {
			result = max(result, dp[0][i]+dp[1][i]+1)
		}
	}

	return result
}

// TLE
func longestMountain11(nums []int) int {
	var (
		l    = len(nums)
		dp   func(prev, curr int, state int, memo [][][3]int) int
		memo = make([][][3]int, l+1)
	)

	for i := 0; i <= l; i++ {
		memo[i] = make([][3]int, l)
		for j := 0; j < l; j++ {
			memo[i][j] = [3]int{-1, -1, -1}
		}
	}

	dp = func(prev, curr int, state int, memo [][][3]int) int {
		if curr == l {
			return 0
		}

		if v := memo[prev+1][curr][state]; v != -1 {
			return v
		}

		dont := dp(prev, curr+1, state, memo)
		do := math.MinInt

		if prev == -1 {
			do = 1 + dp(curr, curr+1, 1, memo)
		} else if state == 1 {
			if nums[prev] < nums[curr] {
				do = 1 + dp(curr, curr+1, 1, memo)
			} else if nums[prev] > nums[curr] {
				do = 1 + dp(curr, curr+1, 2, memo)
			}
		} else if state == 2 {
			if nums[prev] > nums[curr] {
				do = 1 + dp(curr, curr+1, 2, memo)
			}
		}

		if prev == -1 && max(do, dont) < 3 {
			memo[prev+1][curr][state] = 0
		} else {
			memo[prev+1][curr][state] = max(do, dont)
		}

		return memo[prev+1][curr][state]
	}

	return dp(-1, 0, 0, memo)
}
