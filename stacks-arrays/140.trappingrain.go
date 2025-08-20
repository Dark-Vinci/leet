package stacks_arrays

import "slices"

func trap(height []int) int {
	l, r, result := 0, len(height)-1, 0
	lMax, rMax := height[l], height[r]

	for l < r {
		if lMax < rMax {
			l++
			lMax = max(lMax, height[l])
			result += lMax - height[l]
		} else {
			r--
			rMax = max(rMax, height[r])
			result += rMax - height[r]
		}
	}

	return result
}

// this can be sped up by making the height a set and checking if it only contains one val
func trap_EXCEED_TIME(height []int) int {
	l, result := len(height), 0
	maxLeft, maxRight := make([]int, l), make([]int, l)

	for i := 0; i < l; i++ {
		cpLeft, cpRight := slices.Clone(height[:i]), slices.Clone(height[l-i:])
		cpLeft, cpRight = append(cpLeft, 0), append(cpRight, 0)

		maxLeft[i] = slices.Max(cpLeft)
		maxRight[l-i-1] = slices.Max(cpRight)
	}

	for i := 0; i < l; i++ {
		if t := min(maxLeft[i], maxRight[i]) - height[i]; t > 0 {
			result += t
		}
	}

	return result
}
