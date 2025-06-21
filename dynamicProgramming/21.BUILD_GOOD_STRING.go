package dynamicProgramming

func countGoodStrings(low int, high int, zero int, one int) int {
	var helper func(l int, memo map[int]int) int

	helper = func(l int, memo map[int]int) int {
		if l > high {
			return 0
		}

		if v, ok := memo[l]; ok {
			return v
		}

		z, o := helper(l+zero, memo), helper(l+one, memo)

		if l >= low {
			memo[l] = 1 + (z+o)%1000000007
		} else {
			memo[l] = (z + o) % 1000000007
		}

		return memo[l]
	}

	return helper(0, map[int]int{})
}
