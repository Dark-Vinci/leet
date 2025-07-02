package dynamicProgramming

//https://leetcode.com/problems/domino-and-tromino-tiling/description/

func numTilings(n int) int {
	var dp func(x int, memo map[int]int) int

	dp = func(x int, memo map[int]int) int {
		if x < 3 {
			if x == 2 {
				return 2
			}

			return 1
		}

		if v, ok := memo[x]; ok {
			return v
		}

		memo[x] = (2*dp(x-1, memo) + dp(x-3, memo)) % 1000000007

		return memo[x]
	}

	return dp(n, map[int]int{})
}
