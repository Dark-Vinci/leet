package main

import "math"

func maxSubArray(nums []int) int {
	result, maxEnd := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxEnd = max(maxEnd+nums[i], nums[i])
		result = max(result, maxEnd)
	}

	return result
}

// EXCEED TIME LIMIT
func maxSubArray0(nums []int) int {
	result := math.MinInt

	for i := 0; i < len(nums); i++ {
		currSum := 0

		for j := i; j < len(nums); j++ {
			currSum += nums[j]
			result = max(currSum, result)
		}
	}

	return result
}
