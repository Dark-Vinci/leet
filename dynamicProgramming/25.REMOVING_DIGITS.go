package dynamicProgramming

import "math"

//https://cses.fi/problemset/task/1637

func removingDigits(n int) int {
	var dp func(i int, memo map[int]int) int

	dp = func(i int, memo map[int]int) int {
		if i == 0 {
			return 0
		}

		if v, ok := memo[i]; ok {
			return v
		}

		num, result := i, math.MaxInt

		for num != 0 {
			dig := num % 10
			num /= 10

			if dig == 0 {
				continue
			}

			sol := 1 + dp(i-dig, memo)

			result = min(result, sol)
		}

		memo[i] = result

		return memo[i]
	}

	return dp(0, map[int]int{})
}
