package Graphs

func removeStones(stones [][]int) int {
	var (
		result = 0
		u      = NewUnionFind(len(stones))
	)

	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			x0, y0 := stones[i][0], stones[i][1]
			x1, y1 := stones[j][0], stones[j][1]

			if x0 == x1 || y0 == y1 {
				u.Union(i, j)
			}
		}
	}

	for i := 0; i < len(stones); i++ {
		if u.par[i] != i {
			result++
		}
	}

	return result
}
