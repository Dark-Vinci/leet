package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	for i := 0; i < len(dp); i++ {
		if dp[i] == 0 {
			continue
		}

		for _, val := range nums {
			if i+val < len(dp) {
				dp[i+val] += dp[i]
			}
		}
	}

	return dp[target]
}
