package Graphs

import "fmt"

func uniquePaths(m int, n int) int {
	var (
		db  = make(map[string]int)
		dfs func(int, int) int
	)

	dfs = func(right, down int) int {
		if right == 0 && down == 0 {
			return 1
		}

		key := fmt.Sprintf("%d:%d", right, down)

		if v, ok := db[key]; ok {
			return v
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
