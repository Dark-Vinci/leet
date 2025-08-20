package Graphs

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var (
		db  = make(map[string]int)
		dfs func(int, int) int
		n   = len(obstacleGrid)
		m   = len(obstacleGrid[0])
	)

	if obstacleGrid[0][0] == 1 || obstacleGrid[n-1][m-1] == 1 {
		return 0
	}

	dfs = func(right, down int) int {
		if right <= 0 && down <= 0 {
			return 1
		}

		key := fmt.Sprintf("%d:%d", right, down)

		if v, ok := db[key]; ok {
			return v
		}

		if obstacleGrid[n-down-1][m-right-1] == 1 {
			return 0
		}

		l, r := 0, 0

		if right > 0 {
			r = dfs(right-1, down)
		}

		if down > 0 {
			l = dfs(right, down-1)
		}

		db[key] = l + r

		return l + r
	}

	return dfs(m-1, n-1)
}
