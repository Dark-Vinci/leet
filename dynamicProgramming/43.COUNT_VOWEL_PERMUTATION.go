package dynamicProgramming

func countVowelPermutation(n int) int {
	var (
		memo = make([][27]int, n+1)
		dp   func(i int, prev rune, memo [][27]int) int
		MOD  = 1000000007
		VOW  = [5]rune{'a', 'e', 'i', 'o', 'u'}
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

		for _, curr := range VOW {
			if prev == 0 || (prev == 'a' && curr == 'e') || (prev == 'e' && (curr == 'a' || curr == 'i')) || (prev == 'i' && curr != 'i') || (prev == 'o' && (curr == 'i' || curr == 'u')) || (prev == 'u' && curr == 'a') {
				result = (result + dp(nn-1, curr, memo)) % MOD
			}
		}

		memo[nn][key] = result

		return result
	}

	return dp(n, rune(0), memo)
}
