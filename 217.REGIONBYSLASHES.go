package main

func regionsBySlashes(grid []string) int {
	var (
		gridLen = len(grid) * 3
		newGrid = make([][]int, gridLen)
		visited = make(map[[2]int]struct{})
		dfs     func(i, j int)
		result  int
	)

	for i := range newGrid {
		newGrid[i] = make([]int, gridLen)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch {
			case grid[i][j] == '/':
				newGrid[i*3][(j*3)+2] = 1
				newGrid[(i*3)+1][(j*3)+1] = 1
				newGrid[(i*3)+2][(j * 3)] = 1

			case grid[i][j] == '\\':
				newGrid[i*3][(j * 3)] = 1
				newGrid[(i*3)+1][(j*3)+1] = 1
				newGrid[(i*3)+2][(j*3)+2] = 1
			}
		}
	}

	dfs = func(i, j int) {
		if i < 0 || j < 0 || j >= gridLen || i >= gridLen {
			return
		}

		key := [2]int{i, j}

		if _, ok := visited[key]; ok {
			return
		}

		if newGrid[i][j] == 1 {
			return
		}

		visited[key] = struct{}{}

		dfs(i+1, j) // DOWN
		dfs(i-1, j) // UP
		dfs(i, j+1) // RIGHT
		dfs(i, j-1) // LEFT
	}

	for i := range newGrid {
		for j := range newGrid[i] {
			if _, ok := visited[[2]int{i, j}]; !ok && newGrid[i][j] != 1 {
				result++
				dfs(i, j)
			}
		}
	}

	return result
}
