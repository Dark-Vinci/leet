package main

import "slices"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	q := []*TreeNode{root}

	alternator := false

	for len(q) > 0 {
		s := len(q)
		level := make([]int, 0)

		for i := 0; i < s; i++ {
			current := q[0]
			q = q[1:]

			level = append(level, current.Val)

			if current.Left != nil {
				q = append(q, current.Left)
			}

			if current.Right != nil {
				q = append(q, current.Right)
			}
		}

		if alternator {
			slices.Reverse(level)
		}

		result = append(result, level)

		alternator = !alternator
	}

	return result
}
