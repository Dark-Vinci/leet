package Grid

import "math"

func minPathSumTimeLimit(grid [][]int) int {
	var (
		dfs      func(int, int, int)
		h        = len(grid)
		w        = len(grid[0])
		minValue = math.MaxInt
	)

	dfs = func(down, right, sum int) {
		current := grid[h-down-1][w-right-1]

		if down == 0 && right == 0 {
			if current+sum < minValue {
				minValue = current + sum
			}

			return
		}

		if current+sum >= minValue {
			return
		}

		if down > 0 {
			dfs(down-1, right, sum+current)
		}

		if right > 0 {
			dfs(down, right-1, sum+current)
		}
	}

	dfs(h-1, w-1, 0)

	return minValue
}
