package dynamicProgramming

func findLengthOfLCIS(nums []int) int {
	var dp func(curr int, leng int) int

	dp = func(curr int, leng int) int {
		if curr == len(nums) {
			return leng
		}

		if nums[curr] > nums[curr-1] {
			return max(leng, dp(curr+1, leng+1))
		}

		return max(leng, dp(curr+1, 1))
	}

	return dp(1, 1)
}
