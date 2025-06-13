package dynamicProgramming

import "strings"

// RECURSIVE SOLUTION
func findMaxString(strs []string, m, n int) int {
	var (
		dfs  func(i, mm, nn int) int
		memo = make(map[[3]int]int)
	)

	dfs = func(i, mm, nn int) int {
		if i == len(strs) {
			return 0
		}

		key := [3]int{i, mm, nn}

		if v, ok := memo[key]; ok {
			return v
		}

		z, o := strings.Count(strs[i], "0"), strings.Count(strs[i], "1")

		memo[key] = dfs(i+1, mm, nn)

		if z <= mm && n <= nn {
			memo[key] = max(memo[key], 1+dfs(i+1, mm-z, nn-o))
		}

		return memo[key]
	}

	return dfs(0, m, n)
}

func findMaxStrings(strs []string, m, n int) int {
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		z, o := strings.Count(str, "0"), strings.Count(str, "1")

		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = max(
					dp[i][j],
					1+dp[i-z][j-o],
				)
			}
		}
	}

	return dp[m][n]
}
