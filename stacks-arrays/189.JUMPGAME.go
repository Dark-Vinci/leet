package arrays

func canJump(nums []int) bool {
	l := len(nums)
	last := l - 1

	for i := l - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}

	return last == 0
}
