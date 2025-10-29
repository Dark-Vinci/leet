package arrays

import (
	"slices"
)

func maximumImportance(n int, roads [][]int) int64 {
	var (
		connectionCount = make([]int, n)
		result          = int64(0)
		count           = 1
	)

	for _, edg := range roads {
		u, v := edg[0], edg[1]

		connectionCount[u]++
		connectionCount[v]++
	}

	slices.Sort(connectionCount)

	for _, value := range connectionCount {
		result += int64(count) * int64(value)
		count++
	}

	return result
}
