package main

func getMaximumGold(grid [][]int) int {
	var (
		dfs    func(i, j, sum int) int
		cache  = make(map[[2]int]struct{})
		l1, l2 = len(grid), len(grid[0])
		result = 0
	)

	dfs = func(i, j, sum int) int {
		if i >= l2 || j >= l1 || i < 0 || j < 0 {
			return sum
		}

		p := [2]int{i, j}

		if _, ok := cache[p]; ok {
			return 0
		}

		cache[p] = struct{}{}

		if grid[j][i] == 0 {
			return sum
		}

		sum += grid[j][i]

		// UP | DOWN
		up, down := dfs(i, j-1, sum), dfs(i, j+1, sum)

		// RIGHT | LEFT
		right, left := dfs(i+1, j, sum), dfs(i-1, j, sum)

		delete(cache, p)

		return max(up, right, left, down)
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			if grid[j][i] != 0 {
				clear(cache)
				mx := dfs(i, j, 0)
				result = max(result, mx)
			}
		}
	}

	return result
}
