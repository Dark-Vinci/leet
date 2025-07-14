package dynamicProgramming

import (
	"math"
	"slices"
)

func minCost(n int, cuts []int) int {
	var (
		l  = len(cuts) + 2
		dp func(s, e int, memo map[[2]int]int) int
	)

	cuts = append(cuts, 0, n)
	slices.Sort(cuts)

	dp = func(s, e int, memo map[[2]int]int) int {
		if s-e == 1 || s == e || e-s == 1 {
			return 0
		}

		key := [2]int{s, e}
		if v, ok := memo[key]; ok {
			return v
		}

		result := math.MaxInt

		for i := s + 1; i < e; i++ {
			result = min(result, dp(s, i, memo)+dp(i, e, memo))
		}

		memo[key] = result + cuts[e] - cuts[s]

		return memo[key]
	}

	return dp(0, l-1, map[[2]int]int{})
}
