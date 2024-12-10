package main

// NOT ALL TEST CASES PASSED
func countCompleteComponents(n int, edges [][]int) int {
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
