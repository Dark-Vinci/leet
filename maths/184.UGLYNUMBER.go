package maths

func nthUglyNumber(n int) int {
	nums := []int{1}

	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {
		next := min(
			nums[i2]*2,
			nums[i3]*3,
			nums[i5]*5,
		)

		nums = append(nums, next)

		if next == nums[i2]*2 {
			i2++
		}

		if next == nums[i3]*3 {
			i3++
		}

		if next == nums[i5]*5 {
			i5++
		}
	}

	return nums[n-1]
}
