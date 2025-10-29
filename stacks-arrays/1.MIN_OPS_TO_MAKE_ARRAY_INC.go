package arrays

func minOperations(nums []int) int {
	result := 0

	for i := 1; i < len(nums); i++ {
		prev, curr := nums[i-1], nums[i]

		if prev >= curr {
			result += prev - curr + 1
			nums[i] = prev + 1
		}
	}

	return result
}
