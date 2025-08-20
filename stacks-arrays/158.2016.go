package stacks_arrays

import "cmp"

func maximumDifference(nums []int) int {
	mn, mx, diff := nums[0], nums[0], 0

	for i := 1; i < len(nums); i++ {
		prevMn := mn
		mn, mx = min(mn, nums[i]), max(mx, nums[i])

		if prevMn != mn {
			mx = mn
		}

		diff = max(diff, mx-mn)
	}

	return cmp.Or(diff, -1)
}
