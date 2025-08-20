package Graphs

func countUnconnectedComponents(numComponents [][]int) int {
	var (
		result  int
		dfs     func(int)
		visited = make(map[int]struct{})
		adjList = toAdjList(numComponents)
	)

	dfs = func(idx int) {
		if _, ok := visited[idx]; ok {
			return
		}

		visited[idx] = struct{}{}

		for i := 0; i < len(adjList[idx]); i++ {
			dfs(adjList[idx][i])
		}
	}

	for key := range adjList {
		if _, ok := visited[key]; !ok {
			dfs(key)
			result++
		}
	}

	return result
}
