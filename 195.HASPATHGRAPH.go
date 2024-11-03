package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	var (
		list    = toAdjList(edges)
		visited = make(map[int]struct{})
		dfs     func(list map[int][]int, sr int, dt int) bool
	)

	dfs = func(list map[int][]int, sr int, dt int) bool {
		if sr == dt {
			return true
		}

		visited[sr] = struct{}{}
		connected := list[sr]

		for i := 0; i < len(connected); i++ {
			if _, ok := visited[connected[i]]; !ok {
				if dfs(list, connected[i], dt) {
					return true
				}
			}
		}

		return false
	}

	return dfs(list, source, destination)
}

func toAdjList(edges [][]int) map[int][]int {
	result := make(map[int][]int)

	for _, val := range edges {
		u, v := val[0], val[1]

		if _, ok := result[u]; !ok {
			result[u] = []int{}
		}

		if _, ok := result[v]; !ok {
			result[v] = []int{}
		}

		result[u] = append(result[u], v)
		result[v] = append(result[v], u)
	}

	return result
}
