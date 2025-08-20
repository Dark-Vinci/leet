package Grid

func numIslands(grid [][]byte) int {
	type node struct {
		x, y int
	}

	var (
		dfs             func(i, j int)
		l1, l2          = len(grid), len(grid[0])
		result, visited = 0, make(map[node]struct{})
	)

	dfs = func(i, j int) {
		if i >= l2 || j >= l1 || i < 0 || j < 0 {
			return
		}

		if grid[j][i] == '0' {
			return
		}

		n := node{x: i, y: j}

		if _, ok := visited[n]; ok {
			return
		}

		visited[n] = struct{}{}

		dfs(i+1, j) // RIGHT
		dfs(i, j+1) // DOWN
		dfs(i-1, j) // LEFT
		dfs(i, j-1) // UP
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			n := node{x: i, y: j}
			if _, ok := visited[n]; !ok && grid[j][i] == '1' {
				dfs(i, j)
				result++
			}
		}
	}

	return result
}
