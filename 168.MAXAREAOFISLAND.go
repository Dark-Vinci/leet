package main

func maxAreaOfIsland(grid [][]int) int {
	type node struct {
		x, y int
	}

	var (
		l1, l2          = len(grid), len(grid[0])
		visited, result = make(map[node]struct{}), 0
		dfs             func(count *int, i, j int)
	)

	dfs = func(count *int, i, j int) {
		if i >= l2 || j >= l1 || i < 0 || j < 0 {
			return
		}

		n := node{x: i, y: j}

		if grid[j][i] == 0 {
			return
		}

		if _, ok := visited[n]; ok {
			return
		}

		*count += 1

		visited[n] = struct{}{}

		dfs(count, i, j-1) // UP
		dfs(count, i, j+1) // DOWN
		dfs(count, i+1, j) // RIGHT
		dfs(count, i-1, j) // LEFT
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			if grid[j][i] == 1 {
				c := 0
				dfs(&c, i, j)

				result = max(result, c)
			}
		}
	}

	return result
}
