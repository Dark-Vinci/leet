package main

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(r *TreeNode) int

	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		if r.Left == nil && r.Right == nil {
			return 1
		}

		a, b := dfs(r.Left), dfs(r.Right)

		if a == 0 {
			return b + 1
		}

		if b == 0 {
			return a + 1
		}

		return int(math.Min(float64(a), float64(b))) + 1
	}

	return dfs(root)
}
