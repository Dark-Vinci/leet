package main

func shortestPathInGraph(pairs [][]int, s, e int) int {
	type pair struct{ val, dist int }

	var (
		adj = toAdjList(pairs)
		q   = []pair{{s, 0}}
	)

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		if c.val == e {
			return c.dist
		}

		for i := 0; i < len(adj[c.val]); i++ {
			q = append(q, pair{adj[c.val][i], c.dist + 1})
		}
	}

	return -1
}
