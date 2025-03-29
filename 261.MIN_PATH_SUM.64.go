package main

func minPathSum(grid [][]int) int {
	h, w := len(grid), len(grid[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else if j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[h-1][w-1]
}

// time limit
func minPathSum_(grid [][]int) int {
	var (
		dfs func(down, right int) int
		h   = len(grid)
		w   = len(grid[0])
	)

	dfs = func(down, right int) int {
		if down == 0 && right == 0 {
			return grid[down][right]
		}

		if down == 0 {
			return grid[down][right] + dfs(down, right-1)
		}

		if right == 0 {
			return grid[down][right] + dfs(down-1, right)
		}

		return grid[down][right] + min(dfs(down-1, right), dfs(down, right-1))
	}

	return dfs(h-1, w-1)
}
