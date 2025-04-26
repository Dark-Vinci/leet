package dynamicProgramming

func soupServings(n int) float64 {
	if n >= 4800 {
		return 1.0
	}

	var (
		memo = make(map[[2]int]float64)
		dfs  func(a, b int) float64
	)

	n = (n + 24) / 25

	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}

		if a <= 0 {
			return 1.0
		}

		if b <= 0 {
			return 0.0
		}

		key := [2]int{a, b}

		if v, ok := memo[key]; ok {
			return v
		}

		memo[key] = 0.25 * (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))

		return memo[key]
	}

	return dfs(n, n)
}
