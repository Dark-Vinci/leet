package stacks_arrays

import (
	"new_apps/rust_build/melonnu/Trees"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == target {
				return sum
			}

			if Trees.abs(target-sum) < Trees.abs(target-closest) {
				closest = sum
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return closest
}
