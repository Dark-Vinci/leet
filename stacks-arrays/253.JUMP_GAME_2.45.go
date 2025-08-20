package stacks_arrays

import "math"

func jump(nums []int) int {
	var (
		result int
		l      = len(nums)
		i      = 0
	)

	for {
		if i >= l-1 {
			break
		}

		result++

		last := i + 1 + nums[i]

		if last >= l {
			break
		}

		rng := nums[i+1 : last] // I + 1 miss
		mx, mxIndex := math.MinInt, -1

		for j, v := range rng {
			if (v+j) > mx && v > 0 {
				mx = v + j
				mxIndex = j
			}
		}

		i = (i + 1) + mxIndex
	}

	return result
}
