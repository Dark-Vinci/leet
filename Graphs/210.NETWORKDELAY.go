package Graphs

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	type box struct{ to, dist int }

	var (
		list    = make([][]box, n)
		dist    = make([]int, n)
		visited = make(map[int]struct{})
		result  = -1
	)

	for i := range dist {
		dist[i] = math.MaxInt
	}

	for _, val := range times {
		u, v, y := val[0], val[1], val[2]
		list[u-1] = append(list[u-1], box{to: v, dist: y})
	}

	q := [][2]int{{k, 0}}

	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		node, weight := c[0], c[1]

		visited[node-1] = struct{}{}

		if dist[node-1] < weight {
			continue
		}

		dist[node-1] = weight

		for _, v := range list[node-1] {
			if _, ok := visited[v.to-1]; !ok || dist[v.to-1] > v.dist+weight {
				q = append(q, [2]int{v.to, v.dist + weight})
			}
		}
	}

	if len(visited) != n {
		return result
	}

	for _, val := range dist {
		result = max(result, val)
	}

	return result
}
