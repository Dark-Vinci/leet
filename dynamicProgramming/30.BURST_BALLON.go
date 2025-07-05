package dynamicProgramming

// lesson; in place of passing an array as argument, just pass left and right pointer and manipulate as needed
//https://leetcode.com/problems/burst-balloons/description/
// hard

func maxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)

	var dp func(l, r int, memo map[[2]int]int) int

	dp = func(l, r int, memo map[[2]int]int) int {
		if l > r {
			return 0
		}

		key := [2]int{l, r}

		if v, ok := memo[key]; ok {
			return v
		}

		result := 0

		for i := l; i <= r; i++ {
			curr := nums[l-1] * nums[i] * nums[r+1]

			curr += dp(l, i-1, memo) + dp(i+1, r, memo)

			result = max(result, curr)
		}

		memo[key] = result

		return memo[key]
	}

	return dp(1, len(nums)-2, map[[2]int]int{})
}
