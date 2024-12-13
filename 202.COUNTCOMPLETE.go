package main

func countCompleteComponents(n int, edges [][]int) int {
	var (
		result    = 0
		visited   = make(map[int]struct{})
		edgeCount = make(map[int]struct{})
		dfs       func(int)
		list      = toAdjList(edges)
		current   int
	)

	dfs = func(s int) {
		path := list[s]
		edgeCount[len(path)] = struct{}{}
		current += 1

		for _, p := range path {
			if _, ok := visited[p]; !ok {
				visited[p] = struct{}{}
				dfs(p)
			}
		}
	}

	for i := range n {
		clear(edgeCount)
		current = 0

		if _, ok := visited[i]; !ok {
			visited[i] = struct{}{}
			dfs(i)

			if len(edgeCount) == 1 {
				if _, ok := edgeCount[current-1]; ok {
					result++
				}
			}
		}
	}

	return result
}

// NOT ALL TEST CASES PASSED -> APPARENTLY; THIS DOESN;T WORK
func countCompleteComponentsWrong(n int, edges [][]int) int {
	var (
		visited = make(map[int]struct{})
		list    = toAdjList(edges)
		dfs     func(s, prev int) bool
	)

	for i := 0; i < n; i++ {
		if _, ok := list[i]; !ok {
			list[i] = []int{}
		}
	}

	dfs = func(s, prev int) bool {
		if _, ok := visited[s]; ok {
			return false
		}

		visited[s] = struct{}{}

		curr := list[s]

		if len(curr) == 0 {
			return true
		}

		for i := 0; i < len(curr); i++ {
			if curr[i] != prev {
				if !dfs(curr[i], s) {
					return false
				}
			}
		}

		return true
	}

	result := 0

	for key := range list {
		if _, ok := visited[key]; !ok && dfs(key, -1) {
			result++
		}
	}

	return result
}
