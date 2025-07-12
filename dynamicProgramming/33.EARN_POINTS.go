package dynamicProgramming

//https://leetcode.com/problems/number-of-ways-to-earn-points/description/
//difficulty: hard
// lesson
// 1> do, dont can be done in a loop starting from 0 and ending in 1
// when order doesn't matter, loop through the options first

func waysToReachTarget(target int, types [][]int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	for _, typ := range types {
		count, mark := typ[0], typ[1]

		for i := target; i > 0; i-- {
			sum := 0

			for j := 0; j < count; j++ {
				sum += mark

				if sum > i {
					break
				}

				dp[i] = (dp[i] + dp[i-sum]) % 1000000007
			}
		}
	}

	return dp[target]
}

// THIS TLE FOR SOME REASONS UNKNOWN
func waysToReachTarget_(target int, types [][]int) int {
	var dp func(sum int, i int, memo [1001][51]int) int

	dp = func(sum int, i int, memo [1001][51]int) int {
		if sum == target {
			return 1
		}

		if i >= len(types) || sum > target {
			return 0
		}

		if memo[sum][i] != 0 {
			return memo[sum][i]
		}

		count, mark, result := types[i][0], types[i][1], 0

		for j := 0; j <= count; j++ {
			currVal := mark * j

			if sum+currVal <= target {
				result = (result + dp(sum+currVal, i+1, memo)) % 1000000007
			}
		}

		memo[sum][i] = result

		return result
	}

	return dp(0, 0, [1001][51]int{})
}
