package main

import "slices"

func findRedundantConnection(edges [][]int) []int {
	var (
		n       = len(edges)
		dfs     func(lst map[int][]int, visited map[int]struct{}, start int)
		result  []int
		list    map[int][]int
		visited map[int]struct{}
	)

	dfs = func(lst map[int][]int, visited map[int]struct{}, start int) {
		if _, ok := visited[start]; ok {
			return
		}

		visited[start] = struct{}{}

		for i := 0; i < len(lst[start]); i++ {
			dfs(lst, visited, lst[start][i])
		}
	}

	for i := n - 1; i >= 0; i-- {
		clear(visited)

		curr := edges[i]
		edges = slices.Delete(edges, i, i+1)
		list, visited = toAdjList(edges), make(map[int]struct{})

		dfs(list, visited, edges[0][0])

		if len(visited) == n {
			result = curr
			break
		}

		edges = slices.Insert(edges, i, curr)
	}

	return result
}
