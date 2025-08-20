package Graphs

import "slices"

func eventualSafeNodes(graph [][]int) []int {
	var (
		db      = make(map[int]bool)
		result  = make([]int, 0)
		visited = make(map[int]struct{})
		dfs     func(i int) bool
	)

	dfs = func(i int) bool {
		if isSafe, ok := db[i]; ok {
			return isSafe
		}

		if _, ok := visited[i]; ok {
			return false
		}

		visited[i] = struct{}{}
		pos := graph[i]

		if len(pos) == 0 {
			db[i] = true
			return true
		}

		for j := 0; j < len(pos); j++ {
			if !dfs(pos[j]) {
				db[i] = false
				return false
			}
		}

		db[i] = true
		delete(visited, i)

		return true
	}

	for i := range graph {
		if dfs(i) {
			result = append(result, i)
		}
	}

	slices.Sort(result)

	return result
}
