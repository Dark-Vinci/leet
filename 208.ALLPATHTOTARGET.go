package main

import "slices"

func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		result = make([][]int, 0)
		dfs    func(i int, pocket []int)
		list   = make(map[int][]int)
		n      = len(graph)
	)

	for i, val := range graph {
		list[i] = val
	}

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
