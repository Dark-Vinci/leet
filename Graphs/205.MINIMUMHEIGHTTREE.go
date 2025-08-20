package Graphs

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	list, degree := toAdjListDegree(n, edges)
	queue := make([]int, 0)

	for i, d := range degree {
		if d == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		l := len(queue)
		n = n - l

		for i := 0; i < l; i++ {
			v := queue[0]
			queue = queue[1:]

			for _, u := range list[v] {
				degree[u]--

				if degree[u] == 1 {
					queue = append(queue, u)
				}
			}
		}
	}

	return queue
}

func toAdjListDegree(n int, edges [][]int) (map[int][]int, []int) {
	result, degree := make(map[int][]int), make([]int, n)

	for _, tr := range edges {
		u, v := tr[0], tr[1]

		degree[u]++
		degree[v]++

		result[u] = append(result[u], v)
		result[v] = append(result[v], u)
	}

	return result, degree
}
