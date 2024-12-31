package main

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	var (
		dfs     func(at int)
		visited = make(map[int]struct{})
		list    = toAdjList(connections)
		result  = 0
	)

	dfs = func(at int) {
		visited[at] = struct{}{}

		for i := 0; i < len(list[at]); i++ {
			if _, ok := visited[list[at][i]]; !ok {
				dfs(list[at][i])
			}
		}
	}

	for i := range n {
		if _, ok := visited[i]; !ok {
			result++
			dfs(i)
		}
	}

	return result - 1
}
