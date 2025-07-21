package dynamicProgramming

func longestIdealString(s string, k int) int {
	var (
		dp   func(prev, curr int, memo [][27]int) int
		memo = make([][27]int, len(s))
	)

	for i := range memo {
		for j := 0; j < 27; j++ {
			memo[i][j] = -1
		}
	}

	dp = func(prev, curr int, memo [][27]int) int {
		if curr == len(s) {
			return 0
		}

		prevIdx := prev + 1

		if memo[curr][prevIdx] != -1 {
			return memo[curr][prevIdx]
		}

		dont := dp(prev, curr+1, memo)
		do := 0

		currVal := int(s[curr] - 'a')

		if prev == -1 || abs(prev-currVal) <= k {
			do = 1 + dp(currVal, curr+1, memo)
		}

		memo[curr][prevIdx] = max(do, dont)

		return memo[curr][prevIdx]
	}

	return dp(-1, 0, memo)
}
