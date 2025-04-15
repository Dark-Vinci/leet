package dynamicProgramming

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	var (
		dp = make([]int, n)
		mp = make(map[int]int)
	)

	for i := range primes {
		mp[i] = 0
	}

	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt
	}

	dp[0] = 1

	for i := 1; i < n; i++ {
		for j := 0; j < len(primes); j++ {
			dp[i] = min(dp[i], primes[j]*dp[mp[j]])
		}

		for j := 0; j < len(primes); j++ {
			if dp[i] == primes[j]*dp[mp[j]] {
				mp[j]++ // we should break too
			}
		}
	}

	return dp[n-1]
}
