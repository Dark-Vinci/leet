package dynamicProgramming

// GFG
//LINK: https://www.geeksforgeeks.org/dsa/count-number-binary-strings-without-consecutive-1s/

func countStrings(n int) int {
	var helper func(n int, memo map[int]int) int

	helper = func(n int, memo map[int]int) int {
		if n < 3 {
			return n
		}

		if v, ok := memo[n]; ok {
			return v
		}

		memo[n] = helper(n-1, memo) + helper(n-2, memo)

		return memo[n]
	}

	return helper(n, map[int]int{})
}
