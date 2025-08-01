package dynamicProgramming

func countVowelStrings(n int) int {
	var (
		memo = make([][27]int, n+1)
		dp   func(i int, prev rune, memo [][27]int) int
	)

	for i := 0; i <= n; i++ {
		for j := 0; j < 27; j++ {
			memo[i][j] = -1
		}
	}

	dp = func(nn int, prev rune, memo [][27]int) int {
		if nn == 0 {
			return 1
		}

		var key int
		if prev == 0 {
			key = 26
		} else {
			key = int(prev - 'a')
		}

		if v := memo[nn][key]; v != -1 {
			return v
		}

		result := 0

		for a := range 5 {
			if prev == 0 || int(prev-'a') <= a {
				result += dp(nn-1, rune('a'+a), memo)
			}
		}

		memo[nn][key] = result

		return result
	}

	return dp(n, rune(0), memo)
}
