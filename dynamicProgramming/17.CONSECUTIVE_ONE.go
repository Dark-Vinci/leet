package dynamicProgramming

//PLATFORM: Leetcode
// LINK: https://leetcode.com/problems/non-negative-integers-without-consecutive-ones

func findIntegers(n int) int {
	dp := make([]int, 32)

	dp[0], dp[1] = 1, 2

	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	result, prevBit := 0, 0

	for i := 30; i >= 0; i-- {
		if n&(1<<uint(i)) != 0 {
			result += dp[i]

			if prevBit == 1 {
				return result
			}

			prevBit = 1
		} else {
			prevBit = 0
		}
	}

	return result + 1
}
