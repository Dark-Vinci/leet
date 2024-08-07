package main

import "slices"

func merge1(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	result := make([][]int, 0)
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]

		if current[0] <= prev[1] {
			prev[1] = max(prev[1], current[1])
		} else {
			result = append(result, prev)
			prev = current
		}
	}

	// push the last
	result = append(result, prev)

	return result
}
