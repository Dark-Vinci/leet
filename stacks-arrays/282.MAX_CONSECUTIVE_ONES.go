package stacks_arrays

func findMaxConsecutiveOnes(nums []int) int {
	result, count := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result = max(count, result)
			count = 0
			continue
		}

		count++
	}

	return max(result, count)
}
