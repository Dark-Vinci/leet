package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	minValue := math.MaxInt
	var prev *TreeNode

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if prev != nil {
			diff := abs(node.Val - prev.Val)

			if diff < minValue {
				minValue = diff
			}
		}

		prev = node

		dfs(node.Right)
	}

	dfs(root)

	return minValue
}
