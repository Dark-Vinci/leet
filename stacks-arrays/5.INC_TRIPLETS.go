package arrays

import (
	"cmp"
	"math"
	"slices"
)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		if curr <= first {
			first = curr
		} else if curr <= second {
			second = curr
		} else {
			return true
		}
	}

	return false
}

// TLE
func increasingTriplet1(nums []int) bool {
	nums = slices.Compact(nums)

	if sorted := slices.IsSortedFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	}); sorted {
		return false
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[k] > nums[j] && nums[j] > nums[i] {
					return true
				}
			}
		}
	}

	return false
}
