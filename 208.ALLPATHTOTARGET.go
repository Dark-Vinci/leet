package main

import "slices"

func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		n      = len(graph)
		list   = adjList(graph)
		result = make([][]int, 0)
		dfs    func(i int, pocket []int)
	)

	dfs = func(i int, pocket []int) {
		if i == n-1 {
			pocket = append(pocket, i)
			result = append(result, slices.Clone(pocket))
			return
		}

		pocket = append(pocket, i)

		for _, val := range list[i] {
			dfs(val, pocket)
		}
	}

	dfs(0, []int{})

	return result
}

func adjList(graph [][]int) map[int][]int {
	result := make(map[int][]int)

	for i, val := range graph {
		result[i] = val
	}

	return result
}
