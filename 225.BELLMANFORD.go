package main

import (
	"math"
)

func bellmanFord(n int, source int, edges [][]int) []int {
	result := make([]int, n)

	for i := range n {
		result[i] = math.MaxInt
	}

	result[source] = 0

	for range n {
		for _, edj := range edges {
			u, v, w := edj[0], edj[1], edj[2]

			if result[u] != math.MaxInt && result[u]+w < result[v] {
				result[v] = result[u] + w
			}
		}
	}

	return result
}
