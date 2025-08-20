package stacks_arrays

func validMountainArray(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	i, n := 1, len(nums)

	for i < n && nums[i] > nums[i-1] {
		i++
	}

	if i == 1 || i == n {
		return false
	}

	for i < n && nums[i] < nums[i-1] {
		i++
	}

	return i == n
}
