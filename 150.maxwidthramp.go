package main

func maxWidthRamp(nums []int) int {
	maxRight := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			maxRight[i] = nums[i]
			continue
		}

		maxRight[i] = max(nums[i], maxRight[i+1])
	}

	result, left := 0, 0

	for right := range len(nums) {
		for nums[left] > maxRight[right] {
			left++
		}

		result = max(result, right-left)
	}

	return result
}
