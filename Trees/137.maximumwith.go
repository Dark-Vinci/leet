package Trees

import (
	"math"
)

// solution that works for all conditions
func widthOfBinaryTree(root *TreeNode) int {
	type pair struct {
		node  *TreeNode
		index int
	}

	var (
		q      = []pair{{root, 0}}
		result = 0
	)

	for len(q) > 0 {
		l := len(q)
		start := q[0].index
		end := q[l-1].index

		result = max(end-start+1, result)

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]

			if c.node.Left != nil {
				q = append(q, pair{node: c.node.Left, index: 2 * c.index})
			}

			if c.node.Right != nil {
				q = append(q, pair{node: c.node.Right, index: (2 * c.index) + 1})
			}
		}
	}

	return result
}

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
