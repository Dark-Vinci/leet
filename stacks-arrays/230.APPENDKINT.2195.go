package stacks_arrays

import (
	"math"
	"slices"
)

func minimalKSum(nums []int, k int) int64 {
	result := int64(0)

	nums = append(nums, 0, math.MaxInt)

	slices.Sort(nums)

	for i := 0; i < len(nums)-1; i++ {
		change := nums[i+1] - nums[i]

		if change <= 1 {
			continue
		}

		s, start := 0, nums[i]+1

		for start < nums[i+1] {
			if k == 0 {
				break
			}
			k--
			s += start
			start++
		}

		result += int64(s)
	}

	return result
}
