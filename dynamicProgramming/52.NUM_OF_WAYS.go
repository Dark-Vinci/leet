package dynamicProgramming

func numberOfWays2(startPos int, endPos int, k int) int {
	var dp func(curr, count int, memo map[[2]int]int) int

	dp = func(curr, count int, memo map[[2]int]int) int {
		if count == k {
			if curr == endPos {
				return 1
			}

			return 0
		}

		key := [2]int{curr, count}
		if v, ok := memo[key]; ok {
			return v
		}

		minus, plus := dp(curr+1, count+1, memo), dp(curr-1, count+1, memo)

		memo[key] = (minus + plus) % 1000000007

		return memo[key]
	}

	return dp(startPos, 0, map[[2]int]int{})
}
