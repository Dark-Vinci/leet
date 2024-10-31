package main

func search1(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l < r && nums[l] == nums[r] {
		if nums[l] == target || nums[r] == target {
			return true
		}

		l++
		r--
	}

	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target || nums[l] == target || nums[r] == target {
			return true
		}

		if nums[l] <= nums[m] {
			if target >= nums[m] || target <= nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target >= nums[r] || target <= nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return false
}
