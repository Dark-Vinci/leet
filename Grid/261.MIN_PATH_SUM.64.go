package Grid

func minPathSum_(grid [][]int) int {
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

func minPathSum(grid [][]int) int {
	var (
		dfs func(down, right int, memo map[[2]int]int) int
		h   = len(grid)
		w   = len(grid[0])
	)

	dfs = func(down, right int, memo map[[2]int]int) int {
		key := [2]int{down, right}

		if v, ok := memo[key]; ok {
			return v
		}

		if down == 0 && right == 0 {
			memo[key] = grid[down][right]
			return grid[down][right]
		}

		if down == 0 {
			memo[key] = grid[down][right] + dfs(down, right-1, memo)
			return memo[key]
		}

		if right == 0 {
			memo[key] = grid[down][right] + dfs(down-1, right, memo)
			return memo[key]
		}

		memo[key] = grid[down][right] + min(dfs(down-1, right, memo), dfs(down, right-1, memo))
		return memo[key]
	}

	return dfs(h-1, w-1, map[[2]int]int{})
}
