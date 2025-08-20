package stacks_arrays

func triangularSum(nums []int) int {
	for len(nums) > 1 {
		for i := 0; i < len(nums); i++ {
			if i > 0 {
				nums[i-1] += nums[i] % 10
			}
		}

		nums = nums[:len(nums)-1]
	}

	return nums[0] % 10
}
