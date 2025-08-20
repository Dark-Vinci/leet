package Grid

import (
	"cmp"
	"slices"
)

func longestIncreasingPath(matrix [][]int) int {
	type box struct{ i, j, v int }
	type b struct{ i1, j1, i2, j2 int }

	db := make([]box, 0)
	memo := make(map[b]int)

	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[j]); i++ {
			db = append(db, box{i: i, j: j, v: matrix[j][i]})
		}
	}

	slices.SortFunc(db, func(a, b box) int {
		return cmp.Compare(a.v, b.v)
	})

	var dfs func(prev, i, j int) int

	dfs = func(prev, i, j int) int {
		if i >= len(matrix[0]) || j >= len(matrix) || i < 0 || j < 0 {
			return 0
		}

		curr := matrix[j][i]

		if curr <= prev {
			return 0
		}

		item := b{i1: i, j1: j, i2: i, j2: j}

		res := 0

		// up
		item.j2 = j - 1
		if val, ok := memo[item]; ok {
			res = val
		} else {
			res = dfs(curr, i, j-1)
			memo[item] = res
		}
		item.j2 = j

		//down
		item.j2 = j + 1
		if val, ok := memo[item]; ok {
			res = val
		} else {
			res = max(dfs(curr, i, j+1), res)
			memo[item] = res
		}
		item.j2 = j

		//left
		item.i2 = i - 1
		if val, ok := memo[item]; ok {
			res = val
		} else {
			res = max(dfs(curr, i-1, j), res)
			memo[item] = res
		}
		item.i2 = i

		// right
		item.i2 = i + 1
		if val, ok := memo[item]; ok {
			res = val
		} else {
			res = max(dfs(curr, i+1, j), res)
			memo[item] = res
		}

		return 1 + res
	}

	result := 0

	for len(db) > 0 {
		current := db[0]
		db = db[1:]

		result = max(dfs(-1, current.i, current.j), result)
	}

	return result
}
