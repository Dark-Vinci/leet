package dynamicProgramming

func numberOfWays(corridor string) int {
	var (
		dp  func(i int, count int, memo map[[2]int]int) int
		mod = 1_000_000_007
	)

	dp = func(i int, count int, memo map[[2]int]int) int {
		if i >= len(corridor) && count == 2 {
			return 1
		}

		if i >= len(corridor) || count > 2 {
			return 0
		}

		key := [2]int{i, count}
		if v, ok := memo[key]; ok {
			return v
		}

		result := 0

		if corridor[i] != 'S' {
			if count == 2 { // valid
				do, dont := dp(i+1, 0, memo), dp(i+1, count, memo)
				result = (result + do + dont) % mod
			} else if count < 2 {
				result = (result + dp(i+1, count, memo)) % mod
			}
		} else {
			count += 1

			if count == 2 {
				do, dont := dp(i+1, 0, memo), dp(i+1, count, memo) // add barrier and dont add
				result = (result + do + dont) % mod
			} else if count < 2 {
				result = (result + dp(i+1, count, memo)) % mod
			}
		}

		memo[key] = result

		return result
	}

	return dp(0, 0, map[[2]int]int{})
}
