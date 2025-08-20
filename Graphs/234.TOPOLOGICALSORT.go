package Graphs

func topologicalSort(edges [][]int, n int) []int {
	var (
		dfs     func(from int, visited map[int]struct{})
		visited = make(map[int]struct{})
		result  = make([]int, 0)
		list    = meAdjList(edges)
	)

	for i := range n {
		if _, ok := list[i]; !ok {
			list[i] = []int{}
		}
	}

	dfs = func(from int, visited map[int]struct{}) {
		if _, ok := visited[from]; ok {
			return
		}

		visited[from] = struct{}{}

		if len(list[from]) == 0 {
			result = append(result, from)
			return
		}

		for _, to := range list[from] {
			dfs(to, visited)
		}

		result = append([]int{from}, result...)
	}

	for i := range n {
		if _, ok := visited[i]; !ok {
			dfs(i, visited)
		}
	}

	return result
}

func meAdjList(edges [][]int) map[int][]int {
	result := make(map[int][]int)

	for _, edg := range edges {
		u, v := edg[0], edg[1] // u -> v
		result[u] = append(result[u], v)
	}

	return result
}
