package main

import (
	"cmp"
	"slices"
)

func kruskalMst(n int, edges [][]int, force []int) int {
	type edge struct {
		from, to int
		weight   int
	}

	graph := make([]edge, 0)

	for _, edj := range edges {
		u, v, w := edj[0], edj[1], edj[2]

		graph = append(graph, edge{
			from:   u,
			to:     v,
			weight: w,
		})
	}

	slices.SortFunc(graph, func(a, b edge) int {
		return cmp.Compare(a.weight, b.weight)
	})

	if force != nil {
		u, v, w := force[0], force[1], force[2]

		graph = append([]edge{
			{
				from:   u,
				to:     v,
				weight: w,
			},
		}, graph...)
	}

	result, count := 0, 0
	u := NewUnionFind(n)

	for _, edj := range graph {
		if count > n-1 {
			break
		}

		if u.Find(edj.from) != u.Find(edj.to) {
			count++
			result += edj.weight
			u.Union(edj.from, edj.to)
		}
	}

	return result
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	var (
		mn     = kruskalMst(n, edges, nil)
		result = make([][]int, 2)
		l      = len(edges)
	)

	for i := range l {
		removed := edges[i]
		edges = slices.Delete(edges, i, i+1)

		r := kruskalMst(n, edges, removed)

		if r != mn {
			edges = slices.Insert(edges, i, removed)
			continue
		}

		res := kruskalMst(n, edges, nil)

		if res != mn {
			result[0] = append(result[0], i)
		} else {
			result[1] = append(result[1], i)
		}

		edges = slices.Insert(edges, i, removed)
	}

	return result
}
