package Trees

import (
	"math"
)

func findMode(root *TreeNode) []int {
	db := make(map[int]int)

	var dfs func(*TreeNode)

	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		db[r.Val] += 1

		dfs(r.Left)
		dfs(r.Right)
	}

	dfs(root)

	maximum := math.MinInt

	for _, v := range db {
		if v > maximum {
			maximum = v
		}
	}

	result := make([]int, 0)

	for k, v := range db {
		if v == maximum {
			result = append(result, k)
		}
	}

	return result
}
