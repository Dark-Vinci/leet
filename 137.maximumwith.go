package main

import "math"

func widthOfBinaryTreeEXCEED_TIME_LIMIT(root *TreeNode) int {
	var (
		q      = []*TreeNode{root}
		result = 0
	)

	for len(q) > 0 {
		l, start, end := len(q), -1, -1

		for i, j := 0, l-1; i < l && j >= 0; i, j = i+1, j-1 {
			if q[i].Val != math.MaxInt && start == -1 {
				start = i
			}

			if q[j].Val != math.MaxInt && end == -1 {
				end = j
			}
		}

		if start < 0 {
			break
		}

		result = max(end-start+1, result)

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]

			if c.Left != nil {
				q = append(q, c.Left)
			} else {
				q = append(q, &TreeNode{Val: math.MaxInt})
			}

			if c.Right != nil {
				q = append(q, c.Right)
			} else {
				q = append(q, &TreeNode{Val: math.MaxInt})
			}
		}
	}

	return result
}
