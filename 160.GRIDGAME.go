package main

import (
	"math"
	"slices"
)

func gridGame(grid [][]int) int64 {
	prefTop, prefBut := slices.Clone(grid[0]), slices.Clone(grid[1])
	l, result := len(grid[0]), math.MaxInt

	for i := 1; i < l; i++ {
		prefTop[i] += prefTop[i-1]
		prefBut[i] += prefBut[i-1]
	}

	for i := 0; i < l; i++ {
		top, but := prefTop[l-1]-prefTop[i], 0

		if i != 0 {
			but = prefBut[i-1]
		}

		second := max(top, but)
		result = min(result, second)
	}

	return int64(result)
}
