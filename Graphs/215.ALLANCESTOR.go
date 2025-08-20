package Graphs

import "slices"

func getAncestors(n int, edges [][]int) [][]int {
	graph, result := make([][]int, n), make([][]int, n)

	for _, val := range edges {
		graph[val[0]] = append(graph[val[0]], val[1])
	}

	for i := 0; i < n; i++ {
		bfs(i, graph, result)
	}

	return result
}

func bfs(i int, graph [][]int, result [][]int) {
	q, visited := []int{i}, map[int]struct{}{i: {}}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		for _, v := range graph[current] {
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				q = append(q, v)
				result[v] = append(result[v], i)
			}
		}
	}
}

func getAncestors1(n int, edges [][]int) [][]int {
	var (
		visited  = make(map[int]struct{})
		dfs      func(from, to int) bool
		result   = make([][]int, 0)
		list     = make(map[int][]int)
		contains = make(map[int]struct{})
		memo     = make(map[[2]int]bool)
	)

	for _, val := range edges {
		u, v := val[0], val[1] // from -> to

		list[u] = append(list[u], v)
		contains[u], contains[v] = struct{}{}, struct{}{}
	}

	dfs = func(from, to int) bool {
		key := [2]int{from, to}

		if val, ok := memo[key]; ok {
			return val
		}

		if from == to {
			return true
		}

		if _, ok := visited[from]; ok {
			return false
		}

		visited[from] = struct{}{}

		for _, next := range list[from] {
			if dfs(next, to) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false

		return false
	}

	for i := range n {
		if _, ok := contains[i]; !ok {
			result = append(result, []int{})
			continue
		}

		c := make([]int, 0)

		for j := range list {
			if i == j {
				continue
			}

			clear(visited)

			if dfs(j, i) {
				c = append(c, j)
			}
		}

		slices.Sort(c)

		result = append(result, c)
	}

	return result
}
