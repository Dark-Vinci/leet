package main

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

	if diff == 0 {
		return -1
	}

	return diff
}
