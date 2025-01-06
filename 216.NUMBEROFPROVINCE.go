package main

func findCircleNum(isConnected [][]int) int {

	var (
		visited = make(map[int]struct{})
		result  = 0
		adjList = make(map[int][]int)
		dfs     func(to int)
	)

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			if i != j && isConnected[i][j] != 0 {
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	dfs = func(to int) {
		if _, ok := visited[to]; ok {
			return
		}

		visited[to] = struct{}{}

		for i := 0; i < len(adjList[to]); i++ {
			dfs(adjList[to][i])
		}
	}

	for key := range len(isConnected) {
		if _, ok := visited[key]; !ok {
			result++
			dfs(key)
		}
	}

	return result
}
