package dynamicProgramming

import "fmt"

//https://leetcode.com/problems/number-of-ways-to-separate-numbers/

// TLE
func numberOfCombinations(num string) int {
	var (
		dp  func(i int, prev string, memo map[[2]string]int) int
		mod = 1000000007
	)

	dp = func(i int, prev string, memo map[[2]string]int) int {
		if i == len(num) {
			return 1
		}

		if num[i] == '0' {
			return 0
		}

		key := [2]string{fmt.Sprintf("%v", i), prev}

		if v, ok := memo[key]; ok {
			return v
		}

		result := 0

		for j := i + 1; j <= len(num); j++ {
			curr := string(num[i:j])

			if len(prev) < len(curr) || len(prev) == 0 || (len(prev) == len(curr) && curr >= prev) {
				result = (result + dp(j, curr, memo)) % mod
			}
		}

		memo[key] = result

		return memo[key]
	}

	return dp(0, "", map[[2]string]int{})
}
