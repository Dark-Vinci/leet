package main

import (
	"cmp"
	"slices"
)

type point struct {
	x, y, idx int
}

type edge struct {
	from, to point
	weight   int
}

func minCostConnectPoints(points [][]int) int {
	graph := make([]edge, 0)

	for i, a := range points {
		from := point{x: a[0], y: a[1], idx: i}

		for j, b := range points {
			if i != j {
				to := point{x: b[0], y: b[1], idx: j}
				w := calDist(a, b)

				graph = append(graph, edge{
					from:   from,
					to:     to,
					weight: w,
				})
			}
		}
	}

	slices.SortFunc(graph, func(a, b edge) int {
		return cmp.Compare(a.weight, b.weight)
	})

	result, count := 0, 0
	u := NewUnionFind(len(points))

	for _, edg := range graph {
		if count >= len(points) {
			break
		}

		if u.Find(edg.from.idx) != u.Find(edg.to.idx) {
			u.Union(edg.from.idx, edg.to.idx)
			result += edg.weight
			count++
		}
	}

	return result
}

func calDist(a, b []int) int {
	xa, xb, ya, yb := a[0], b[0], a[1], b[1]

	x, y := abs(xa-xb), abs(ya-yb)

	return x + y
}
