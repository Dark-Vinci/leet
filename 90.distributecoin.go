package main

import "math"

func distributeCoins(root *TreeNode) int {
	var (
		result int
		dfs    func(*TreeNode) (int, int)
	)

	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		l, leftWeight := dfs(root.Left)
		r, rightWeight := dfs(root.Right)

		c := l + r + 1
		cWeight := leftWeight + rightWeight + root.Val

		result += int(math.Abs(float64(c - cWeight)))

		return c, cWeight
	}

	dfs(root)

	return result
}
