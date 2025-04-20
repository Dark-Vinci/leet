package main

import "slices"

func countFairPairs(nums []int, lower int, upper int) int64 {
	var (
		result = int64(0)
		bin    func(i int, target int, up bool) int
	)

	slices.Sort(nums)

	bin = func(i int, target int, up bool) int {
		l, r := i, len(nums)-1

		for l <= r {
			m := l + (r-l)/2

			if up {
				if nums[m] <= target {
					l = m + 1
				} else {
					r = m - 1
				}
			} else {
				if nums[m] < target {
					l = m + 1
				} else {
					r = m - 1
				}
			}
		}

		return l
	}

	for i := 0; i < len(nums); i++ {
		u := bin(i+1, upper-nums[i], true)
		l := bin(i+1, lower-nums[i], false)

		result += int64(u - l)
	}

	return result
}

func countFairPairs_(nums []int, lower int, upper int) int64 {
	result := int64(0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if isValid__(lower, upper, nums[i], nums[j]) {
				result += 1
			}
		}
	}

	return result
}

func isValid__(l, u, a, b int) bool {
	return a+b >= l && a+b <= u
}
