package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var (
		result   = make([]float64, len(queries))
		db       = make(map[[2]string]float64)
		list     = toAdjList(equations)
		visited  = make(map[string]struct{})
		contains = make(map[string]struct{})
		dfs      func(visited map[string]struct{}, from, to string) (float64, bool)
	)

	for i := range len(values) {
		a, b, q := equations[i][0], equations[i][1], values[i] // from , to , val
		db[[2]string{a, b}] = q
		db[[2]string{b, a}] = 1 / q
		contains[a] = struct{}{}
		contains[b] = struct{}{}
	}

	dfs = func(visited map[string]struct{}, from, to string) (float64, bool) {
		if _, ok := contains[from]; !ok {
			return -1.0, false
		}

		if _, ok := contains[to]; !ok {
			return -1.0, false
		}

		if from == to {
			return 1.0, true
		}

		if _, ok := visited[from]; ok {
			return 1.0, false
		}

		visited[from] = struct{}{}

		for i := 0; i < len(list[from]); i++ {
			if val, ok := dfs(visited, list[from][i], to); ok {
				return val * db[[2]string{from, list[from][i]}], true
			}
		}

		return -1.0, false
	}

	for i, val := range queries {
		clear(visited)

		u, v := val[0], val[1] // from -> to

		res, _ := dfs(visited, u, v)

		result[i] = res
	}

	return result
}
