package arrays

import (
	"math"
	"slices"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	slices.Sort(nums)

	maxDiff := math.MinInt

	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		diff := int(math.Abs(float64(nums[i] - nums[j])))

		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}
