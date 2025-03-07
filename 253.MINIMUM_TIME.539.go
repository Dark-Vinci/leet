package main

import (
	"math"
	"slices"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	minutes := make([]int, len(timePoints))

	for i, time := range timePoints {
		h, m := time[0:2], time[3:5]

		hInt, _ := strconv.Atoi(h)
		mInt, _ := strconv.Atoi(m)

		a := hInt*60 + mInt

		minutes[i] = a
	}

	slices.Sort(minutes)

	result := math.MaxInt

	for i := 1; i < len(minutes); i++ {
		a, b := minutes[i-1], minutes[i]

		result = min(b-a, result)
	}

	return min(result, abs(minutes[0]+1440-minutes[len(minutes)-1]))
}
