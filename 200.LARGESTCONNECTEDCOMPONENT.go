package main

func largestConnectedComponent(components [][]int) int {
	var (
		adjList = toAdjList(components)
		result  int
		dfs     func(int) int
		visited = make(map[int]struct{})
	)

	dfs = func(i int) int {
		if _, ok := visited[i]; ok {
			return 0
		}
		visited[i] = struct{}{}

		res := 1

		for j := 0; j < len(adjList[i]); j++ {
			res += dfs(adjList[i][j])
		}

		return res
	}

	for key := range adjList {
		if _, ok := visited[key]; !ok {
			result = max(result, dfs(key))
		}
	}

	return result
}
